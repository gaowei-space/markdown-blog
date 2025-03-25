package app

import (
	"fmt"
	"html/template"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/gaowei-space/markdown-blog/internal/api"
	"github.com/gaowei-space/markdown-blog/internal/bindata/assets"
	"github.com/gaowei-space/markdown-blog/internal/bindata/views"
	"github.com/gaowei-space/markdown-blog/internal/types"
	"github.com/gaowei-space/markdown-blog/internal/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/kataras/iris/v12/view"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"github.com/urfave/cli/v2"
)

var (
	MdDir      string // 指定markdown文件所在目录
	Env        string // 指定环境，区分开发环境（dev）和生产环境（prod）
	Title      string // 指定博客标题
	Index      string // 指定首页路径
	ICP        string // 指定ICP备案号
	ISF        string // 指定公安备案号
	FDir       string // 指定文件目录
	Copyright  int64 // 指定博客版权年份
	LayoutFile = "layouts/layout.html" // 指定HTML布局文件路径
	LogsDir    = "cache/logs/" // 指定日志文件所在目录
	TocPrefix  = "[toc]" // 指定TOC前缀
	IgnoreFile = []string{`favicon.ico`, `.DS_Store`, `.gitignore`, `README.md`} // 指定忽略文件
	IgnorePath = []string{`.git`} // 指定忽略路径
	Cache      time.Duration // 指定页面缓存时间
	Analyzer   types.Analyzer // 指定分析器
	Gitalk     types.Gitalk // 指定Gitalk
)

// web服务器默认端口
const DefaultPort = 5006

func RunWeb(ctx *cli.Context) error {
	initParams(ctx) // 初始化参数

	app := iris.New() // 创建新的Iris应用

	setLog(app) // 设置日志

	app.RegisterView(getTmpl()) // 注册视图
	app.OnErrorCode(iris.StatusNotFound, api.NotFound) // 设置404错误处理
	app.OnErrorCode(iris.StatusInternalServerError, api.InternalServerError) // 设置500错误处理

	setIndexAuto := false // 是否自动设置首页
	if Index == "" {
		setIndexAuto = true
	}

	app.Use(func(ctx iris.Context) {
		activeNav := getActiveNav(ctx) // 获取当前活跃导航项

		navs, firstNav := getNavs(activeNav)// 生成完整的导航菜单，并标记当前活跃导航项

		firstLink := utils.CustomURLEncode(strings.TrimPrefix(firstNav.Link, "/"))
		if setIndexAuto && Index != firstLink {
			Index = firstLink // 自动设置首页路径
		}

		// 设置 Gitalk ID
		Gitalk.Id = utils.MD5(activeNav) // 为评论系统生成唯一 ID

		// 设置视图数据
		ctx.ViewData("Gitalk", Gitalk)
		ctx.ViewData("Analyzer", Analyzer)
		ctx.ViewData("Title", Title)
		ctx.ViewData("Nav", navs)
		ctx.ViewData("ICP", ICP)
		ctx.ViewData("ISF", ISF)
		ctx.ViewData("Copyright", Copyright)
		ctx.ViewData("ActiveNav", activeNav)
		ctx.ViewLayout(LayoutFile) // 设置布局模板

		ctx.Next() // 继续处理请求
	})

	app.Favicon("./favicon.ico") // 配置网站图标
	app.HandleDir("/static", getStatic()) // 配置静态资源文件夹
	app.Get("/{f:path}", iris.Cache(Cache), articleHandler) // 渲染 Markdown 文件为 HTML
	app.Get(fmt.Sprintf("/%s/{f:path}", FDir), serveFileHandler)  // 提供文件下载服务
	app.Run(iris.Addr(":" + strconv.Itoa(parsePort(ctx)))) //启动 Iris Web 应用

	return nil
}

// 获取静态文件目录
func getStatic() interface{} {
	if Env == "prod" { // 生产环境
		return assets.AssetFile()
	} else { // 开发环境
		return "./web/assets"
	}
}

// 获取模板
func getTmpl() *view.HTMLEngine {
	if Env == "prod" { // 生产环境production
		return iris.HTML(views.AssetFile(), ".html").Reload(false)
	} else { // 开发环境development
		return iris.HTML("./web/views", ".html").Reload(true)
	}
}

// 初始化参数
func initParams(ctx *cli.Context) {
	MdDir = ctx.String("dir") // 获取用户传入的--dir参数，指定Markdown文件存放目录
	if strings.TrimSpace(MdDir) == "" {
		log.Panic("Markdown files folder cannot be empty")
	}
	MdDir, _ = filepath.Abs(MdDir) // 将MdDir转换为绝对路径

	Env = ctx.String("env") // 获取用户传入的--env参数，指定环境，区分开发环境（dev）和生产环境（prod）
	Title = ctx.String("title") // 获取用户传入的--title参数，指定博客标题
	Index = ctx.String("index") // 获取用户传入的--index参数，指定首页路径
	ICP = ctx.String("icp") // 获取用户传入的--icp参数，指定ICP备案号
	ISF = ctx.String("isf") // 获取用户传入的--isf参数，指定公安备案号
	Copyright = ctx.Int64("copyright") // 获取用户传入的--copyright参数，指定博客版权年份
	FDir = ctx.String("fdir") // 获取用户传入的--fdir参数，指定文件目录

	Cache = time.Minute * 0 // 设置页面缓存时间为0分钟，即不缓存
	if Env == "prod" {
		Cache = time.Minute * time.Duration(ctx.Int64("cache")) // 如果环境为生产环境，则使用用户传入的--cache参数，指定页面缓存时间
	}

	// 设置分析器
	Analyzer.SetAnalyzer(ctx.String("analyzer-baidu"), ctx.String("analyzer-google"))

	// 设置Gitalk 
	Gitalk.SetGitalk(ctx.String("gitalk.client-id"), ctx.String("gitalk.client-secret"), ctx.String("gitalk.repo"), ctx.String("gitalk.owner"), ctx.StringSlice("gitalk.admin"), ctx.StringSlice("gitalk.labels"))

	// 忽略文件
	IgnoreFile = append(IgnoreFile, ctx.StringSlice("ignore-file")...)
	IgnorePath = append(IgnorePath, FDir)
	IgnorePath = append(IgnorePath, ctx.StringSlice("ignore-path")...)
}

// 设置日志
func setLog(app *iris.Application) {
	os.MkdirAll(LogsDir, 0777) // 确保日志目录存在，如果目录不存在，则创建目录
	// 打开日志文件，如果文件不存在，则创建文件，如果文件存在，则追加写入(access-YYYYMMDD.log)
	f, _ := os.OpenFile(LogsDir+"access-"+time.Now().Format("20060102")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)

	if Env == "prod" {
		app.Logger().SetOutput(f) // 如果环境为生产环境，则将日志输出到文件
	} else {
		app.Logger().SetLevel("debug") // 如果环境为开发环境，则将日志级别设置为debug
		app.Logger().Debugf(`Log level set to "debug"`) // 日志输出到标准输出（如终端）
	}

	// Close the file on shutdown.
	// 注册一个程序关闭时执行的回调函数，关闭日志文件
	app.ConfigureHost(func(su *iris.Supervisor) {
		su.RegisterOnShutdown(func() {
			f.Close()
		})
	})

	// 创建一个accesslog实例，将http请求日志写入到文件
	ac := accesslog.New(f)
	ac.AddOutput(app.Logger().Printer) // 将请求日志同时输出到日志文件和终端
	app.UseRouter(ac.Handler) // 使用accesslog中间件
	app.Logger().Debugf("Using <%s> to log requests", f.Name()) // 记录日志文件路径
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

// 从markdown文件目录结构生成导航菜单，标记当前活跃的导航项，找到第一个非目录的导航项作为首页的默认选项
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

// 递归遍历node树来找到当前活跃导航项
func searchActiveNav(node *utils.Node, activeNav string) {
	link_str, _ := url.QueryUnescape(node.Link)
	if !node.IsDir && strings.TrimPrefix(link_str, "/") == activeNav {
		node.Active = "active"
		return
	}
	if len(node.Children) > 0 {
		for _, v := range node.Children {
			searchActiveNav(v, activeNav)
		}
	}
}

//找到第一个非目录文件
func getFirstNav(node utils.Node) utils.Node {
	if !node.IsDir {
		return node
	}
	return getFirstNav(*node.Children[0])
}

// 获取当前活跃导航项
func getActiveNav(ctx iris.Context) string {
	f := ctx.Params().Get("f")
	if f == "" {
		f = Index
	}
	return f
}

// 提供文件下载
func serveFileHandler(ctx iris.Context) {
	f := ctx.Params().Get("f")
	file := MdDir + "/" + FDir + "/" + f
	ctx.ServeFile(file)
}

// 渲染Markdown文件为HTML
func articleHandler(ctx iris.Context) {
	f := getActiveNav(ctx)

	if utils.IsInSlice(IgnoreFile, f) {
		return
	}

	mdfile := MdDir + "/" + f + ".md"

	// 检查文件是否存在
	_, err := os.Stat(mdfile)
	if err != nil {
		ctx.StatusCode(404)
		ctx.Application().Logger().Errorf("Not Found '%s', Path is %s", mdfile, ctx.Path())
		return
	}

	// 读取文件内容
	bytes, err := os.ReadFile(mdfile)
	if err != nil {
		ctx.StatusCode(500)
		ctx.Application().Logger().Errorf("ReadFile Error '%s', Path is %s", mdfile, ctx.Path())
		return
	}
	tmp := strings.Split(f, "/")
	title := tmp[len(tmp)-1]
	ctx.ViewData("Title", title+" - "+Title)
	ctx.ViewData("Article", mdToHtml(bytes))

	ctx.View("index.html")
}

// 将Markdown文件转换为HTML
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

	// fix windows \r\n
	unix := strings.ReplaceAll(strs, "\r\n", "\n")

	unsafe := blackfriday.Run([]byte(unix), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(blackfriday.CommonExtensions))

	// 创建bluemonday策略，只允许<span>标签及其style属性
	p := bluemonday.UGCPolicy()
	p.AllowElements("span")                  // 只允许<span>标签
	p.AllowAttrs("style").OnElements("span") // 在<span>上允许使用style属性

	// 使用自定义的bluemonday策略来清理HTML
	html := p.SanitizeBytes(unsafe)

	return template.HTML(string(html))
}