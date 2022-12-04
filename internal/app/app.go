package app

import (
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gaowei-space/markdown-blog/internal/types"
	"github.com/urfave/cli/v2"
)

var (
	Port       string
	MdDir      string
	Env        string
	Title      string
	Index      string
	LayoutFile = "layouts/layout.html"
	LogsDir    = "cache/logs/"
	TocPrefix  = "[toc]"
	IgnoreFile = []string{`favicon.ico`, `.DS_Store`, `.gitignore`, `README.md`}
	IgnorePath = []string{`.git`}
	Cache      time.Duration
	Analyzer   types.Analyzer
	Gitalk     types.Gitalk
	Listener   types.Listener
)

// web服务器默认端口
const DefaultPort = 5006

func Run(ctx *cli.Context) error {
	initParams(ctx)

	go RunListener(ctx)

	RunWeb(ctx)

	return nil
}

func initParams(ctx *cli.Context) {
	MdDir = ctx.String("dir")
	if strings.TrimSpace(MdDir) == "" {
		log.Panic("Markdown files folder cannot be empty")
	}
	MdDir, _ = filepath.Abs(MdDir)

	Env = ctx.String("env")
	Title = ctx.String("title")
	Index = ctx.String("index")
	Port = strconv.Itoa(parsePort(ctx))

	Cache = time.Minute * 0
	if Env == "prod" {
		Cache = time.Minute * 3
	}

	// 设置分析器
	Analyzer.SetAnalyzer(ctx.String("analyzer-baidu"), ctx.String("analyzer-google"))

	// 设置Gitalk
	Gitalk.SetGitalk(ctx.String("gitalk.client-id"), ctx.String("gitalk.client-secret"), ctx.String("gitalk.repo"), ctx.String("gitalk.owner"), ctx.StringSlice("gitalk.admin"), ctx.StringSlice("gitalk.labels"))

	// 忽略文件
	IgnoreFile = append(IgnoreFile, ctx.StringSlice("ignore-file")...)
	IgnorePath = append(IgnorePath, ctx.StringSlice("ignore-path")...)

	// 设置 Listener
	Listener.Cert.KeyFile = ctx.String("listener.key-file")
	Listener.Cert.CertFile = ctx.String("listener.cert-file")
	Listener.SetListener(ctx.Bool("listener.open"), ctx.String("listenter.host"), Listener.Cert)
}
