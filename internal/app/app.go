package app

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/gaowei-space/markdown-blog/internal/types"
	"github.com/gaowei-space/markdown-blog/internal/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"github.com/urfave/cli"
)

var (
	MdDir       string
	Env         string
	Title       string
	Index       string
	LayoutFile  = "web/views/layouts/layout.html"
	ArticlesDir = "cache/articles/"
	LogsDir     = "cache/logs/"
	AssetsDir   = "web/assets"
	TocPrefix   = "[toc]"
	Cache       time.Duration
	Analyzer    types.Analyzer
)

// web服务器默认端口
const DefaultPort = 5006

func initParams(ctx *cli.Context) {
	MdDir = ctx.String("dir")
	if strings.TrimSpace(MdDir) == "" {
		log.Panic("Markdown files folder cannot be empty")
	}
	MdDir, _ = filepath.Abs(MdDir)

	Env = ctx.String("env")
	Title = ctx.String("title")
	Index = ctx.String("index")

	Cache = time.Minute * 0
	if Env == "prod" {
		Cache = time.Minute * 3
	}

	// 设置分析器
	Analyzer.SetAnalyzer(ctx.String("analyzer-baidu"), ctx.String("analyzer-google"))
}

func setLog(app *iris.Application) {
	os.MkdirAll(LogsDir, 0777)
	f, _ := os.OpenFile(LogsDir+"access-"+time.Now().Format("20060102")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)

	if Env == "prod" {
		app.Logger().SetOutput(f)
	} else {
		app.Logger().SetLevel("debug")
		app.Logger().Debugf(`Log level set to "debug"`)
	}

	// Close the file on shutdown.
	app.ConfigureHost(func(su *iris.Supervisor) {
		su.RegisterOnShutdown(func() {
			f.Close()
		})
	})

	ac := accesslog.New(f)
	ac.AddOutput(app.Logger().Printer)
	app.UseRouter(ac.Handler)
	app.Logger().Debugf("Using <%s> to log requests", f.Name())
}

func RunWeb(ctx *cli.Context) {
	initParams(ctx)
	app := iris.New()

	setLog(app)

	tmpl := iris.HTML("./", ".html").Reload(true)
	app.RegisterView(tmpl)
	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)

	setIndexAuto := false
	if Index == "" {
		setIndexAuto = true
	}

	app.Use(func(ctx iris.Context) {
		navs, firstNav := getNavs(getActiveNav(ctx))

		firstLink := strings.TrimPrefix(firstNav.Link, "/")
		if setIndexAuto && Index != firstLink {
			Index = firstLink
		}

		ctx.ViewData("Analyzer", Analyzer)
		ctx.ViewData("Title", Title)
		ctx.ViewData("Nav", navs)
		ctx.ViewData("ActiveNav", getActiveNav(ctx))
		ctx.ViewLayout(LayoutFile)

		ctx.Next()
	})

	app.HandleDir("/static", AssetsDir)

	app.Get("/{f:path}", iris.Cache(Cache), show)

	app.Run(iris.Addr(":" + strconv.Itoa(parsePort(ctx))))
}

func parsePort(ctx *cli.Context) int {
	port := DefaultPort
	if ctx.IsSet("port") {
		port = ctx.Int("port")
	}
	if port <= 0 || port >= 65535 {
		port = DefaultPort
	}

	return port
}

func getNavs(activeNav string) ([]map[string]interface{}, utils.Node) {
	var option utils.Option
	option.RootPath = []string{MdDir}
	option.SubFlag = true
	option.IgnorePath = []string{`.git`}
	option.IgnoreFile = []string{`.DS_Store`, `.gitignore`, `README.md`}
	tree, _ := utils.Explorer(option)

	navs := make([]map[string]interface{}, 0)
	for _, v := range tree.Children {
		for _, item := range v.Children {
			searchActiveNav(item, activeNav)
			navs = append(navs, structs.Map(item))
		}
	}

	firstNav := getFirstNav(*tree.Children[0])

	return navs, firstNav
}

func searchActiveNav(node *utils.Node, activeNav string) {
	if !node.IsDir && node.Link == "/"+activeNav {
		node.Active = "active"
		return
	}
	if len(node.Children) > 0 {
		for _, v := range node.Children {
			searchActiveNav(v, activeNav)
		}
	}
}

func getFirstNav(node utils.Node) utils.Node {
	if !node.IsDir {
		return node
	}
	return getFirstNav(*node.Children[0])
}

func getActiveNav(ctx iris.Context) string {
	f := ctx.Params().Get("f")
	if f == "" {
		f = Index
	}
	return f
}

func show(ctx iris.Context) {
	f := getActiveNav(ctx)
	mdfile := MdDir + "/" + f + ".md"
	articlefile := ArticlesDir + f + ".html"

	_, err := os.Stat(mdfile)
	if err != nil {
		ctx.StatusCode(404)
		ctx.Application().Logger().Errorf("Not Found '%s', Path is %s", mdfile, ctx.Path())
		return
	}

	bytes, err := os.ReadFile(mdfile)
	if err != nil {
		ctx.StatusCode(500)
		ctx.Application().Logger().Errorf("ReadFile Error '%s', Path is %s", mdfile, ctx.Path())
		return
	}

	if err := mdToHtml(bytes, articlefile); err != nil {
		ctx.StatusCode(500)
		ctx.Application().Logger().Errorf("WriteFile Error %s, Path is %s", err, ctx.Path())
		return
	}

	ctx.View(articlefile)
}

func mdToHtml(content []byte, filename string) error {
	os.MkdirAll(filepath.Dir(filename), 0777)

	strs := string(content)

	var htmlFlags blackfriday.HTMLFlags

	if strings.HasPrefix(strs, TocPrefix) {
		htmlFlags |= blackfriday.TOC
		strs = strings.Replace(strs, TocPrefix, "<br/><br/>", 1)
	}

	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		Flags: htmlFlags,
	})

	unsafe := blackfriday.Run([]byte(strs), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(blackfriday.CommonExtensions))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	if err := os.WriteFile(filename, html, 0777); err != nil {
		return err
	}

	return nil
}

func notFound(ctx iris.Context) {
	ctx.View("web/views/errors/404.html")
}

func internalServerError(ctx iris.Context) {
	ctx.View("web/views/errors/500.html")
}
