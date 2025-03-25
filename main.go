package main

import (
	"os"
	"strconv"
	"time"

	"github.com/gaowei-space/markdown-blog/internal/app"
	"github.com/gaowei-space/markdown-blog/internal/utils"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

//go:generate go-bindata -fs -o internal/bindata/views/views.go -pkg=views -prefix=web/views ./web/views/...
//go:generate go-bindata -fs -o internal/bindata/assets/assets.go -pkg=assets -prefix=web/assets ./web/assets/...

var (
	MdDir                = "md/"
	Title                = "Blog"
	AppVersion           = "1.1.1"
	BuildDate, GitCommit string
)

// web服务器默认端口
const DefaultPort = 5006

func main() {
	cliApp := cli.NewApp()
	cliApp.Name = "markdown-blog"
	cliApp.Usage = "Markdown Blog App"
	cliApp.Version, _ = utils.FormatAppVersion(AppVersion, GitCommit, BuildDate)
	cliApp.Commands = getCommands()
	cliApp.Flags = append(cliApp.Flags, []cli.Flag{}...) // 没有定义全局标志，这里为空
	cliApp.Run(os.Args)
}

// 获取命令
func getCommands() []*cli.Command {
	web := webCommand()

	return []*cli.Command{web}
}

// 获取web命令
func webCommand() *cli.Command {
	// 通用命令
	commonFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "config", // 标志的全名 --config 
			Value: "", // 标志的默认值
			Usage: "Load configuration from `FILE`, default is empty", // 标志的描述
		},
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "dir", // 标志的全名 --dir
			Aliases: []string{"d"}, // 标志的别名 -d
			Value:   MdDir, // 标志的默认值
			Usage:   "Markdown files dir", // 标志的描述
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "title",
			Aliases: []string{"t"},
			Value:   Title,
			Usage:   "Blog title",
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Value:   DefaultPort,
			Usage:   "Bind port",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "env",
			Aliases: []string{"e"},
			Value:   "prod",
			Usage:   "Runtime environment, dev|test|prod",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "index",
			Aliases: []string{"i"},
			Value:   "",
			Usage:   "Home page, default is empty",
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:    "cache",
			Aliases: []string{"c"},
			Value:   3,
			Usage:   "The cache time unit is minutes, this parameter takes effect in the prod environment",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "icp",
			Value: "",
			Usage: "ICP, default is empty",
		}), altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "isf",
			Value: "",
			Usage: "National Internet Security Filing, default is empty",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "copyright",
			Value: strconv.Itoa(time.Now().Year()),
			Usage: "Copyright, default the current year, such as 2024",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "fdir",
			Value: "public",
			Usage: "File directory name",
		}),
	}

	// gitalk评论系统命令
	gitalkFlags := []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "gitalk.client-id",
			Usage: "Set up Gitalk ClientId, default is empty",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "gitalk.client-secret",
			Usage: "Set up Gitalk ClientSecret, default is empty",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "gitalk.repo",
			Usage: "Set up Gitalk Repo, default is empty",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "gitalk.owner",
			Usage: "Set up Gitalk Repo, default is empty",
		}),
		altsrc.NewStringSliceFlag(&cli.StringSliceFlag{
			Name:  "gitalk.admin",
			Usage: "Set up Gitalk Admin, default is `[gitalk.owner]`",
		}),
		altsrc.NewStringSliceFlag(&cli.StringSliceFlag{
			Name:  "gitalk.labels",
			Usage: "Set up Gitalk Admin, default is `[\"gitalk\"]`",
		}),
	}

	flags := append(commonFlags, gitalkFlags...)

	// 分析器命令，百度和谷歌
	analyzerFlags := []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "analyzer-baidu",
			Aliases: []string{"ab"},
			Value:   "",
			Usage:   "Set up Baidu Analyzer, default is empty",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "analyzer-google",
			Aliases: []string{"ag"},
			Value:   "",
			Usage:   "Set up Google Analyzer, default is empty",
		}),
	}

	flags = append(flags, analyzerFlags...)

	// 忽略文件和路径命令
	ignoreFlags := []cli.Flag{
		altsrc.NewStringSliceFlag(&cli.StringSliceFlag{
			Name:  "ignore-file",
			Usage: "Set up ignore file, eg: demo.md",
		}),
		altsrc.NewStringSliceFlag(&cli.StringSliceFlag{
			Name:  "ignore-path",
			Usage: "Set up ignore path, eg: demo",
		}),
	}

	flags = append(flags, ignoreFlags...)

	web := cli.Command{
		Name:   "web",
		Usage:  "Run blog web server",
		Action: app.RunWeb,
		Flags:  flags,
		Before: altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config")),
		// 在运行命令之前，从配置文件中读取标志，并将读取的标志绑定到命令的标志上
	}

	return &web
}
