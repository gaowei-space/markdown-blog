package app

import (
	"html/template"
	"os"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/gaowei-space/markdown-blog/internal/api"
	"github.com/gaowei-space/markdown-blog/internal/bindata/assets"
	"github.com/gaowei-space/markdown-blog/internal/bindata/views"
	"github.com/gaowei-space/markdown-blog/internal/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"github.com/urfave/cli/v2"
)

func RunWeb(ctx *cli.Context) {
	app := iris.New()

	setLog(app)

	tmpl := iris.HTML(views.AssetFile(), ".html").Reload(true)
	app.RegisterView(tmpl)
	app.OnErrorCode(iris.StatusNotFound, api.NotFound)
	app.OnErrorCode(iris.StatusInternalServerError, api.InternalServerError)

	setIndexAuto := false
	if Index == "" {
		setIndexAuto = true
	}

	app.Use(func(ctx iris.Context) {
		activeNav := getActiveNav(ctx)

		navs, firstNav := getNavs(activeNav)

		firstLink := strings.TrimPrefix(firstNav.Link, "/")
		if setIndexAuto && Index != firstLink {
			Index = firstLink
		}

		// 设置 Gitalk ID
		Gitalk.Id = utils.MD5(activeNav)

		ctx.ViewData("Gitalk", Gitalk)
		ctx.ViewData("Analyzer", Analyzer)
		ctx.ViewData("Title", Title)
		ctx.ViewData("Nav", navs)
		ctx.ViewData("ActiveNav", activeNav)
		ctx.ViewLayout(LayoutFile)

		ctx.Next()
	})

	app.Favicon("./favicon.ico")
	app.HandleDir("/static", assets.AssetFile())
	app.Get("/{f:path}", iris.Cache(Cache), articleHandler)

	app.Run(iris.Addr(":" + Port))
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
	option.IgnorePath = IgnorePath
	option.IgnoreFile = IgnoreFile
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

func articleHandler(ctx iris.Context) {
	f := getActiveNav(ctx)

	if utils.IsInSlice(IgnoreFile, f) {
		return
	}

	mdfile := MdDir + "/" + f + ".md"

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

	ctx.ViewData("Article", mdToHtml(bytes))

	ctx.View("index.html")
}

func mdToHtml(content []byte) template.HTML {
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

	return template.HTML(string(html))
}
