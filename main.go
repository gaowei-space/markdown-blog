package main

import (
	"os"

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
	AppVersion           = "1.0.0"
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
	cliApp.Flags = append(cliApp.Flags, []cli.Flag{}...)
	cliApp.Run(os.Args)
}

func getCommands() []*cli.Command {
	web := webCommand()

	return []*cli.Command{web}
}

func webCommand() *cli.Command {
	commonFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "config",
			Value: "",
			Usage: "Load configuration from `FILE`, default is empty",
		},
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "dir",
			Aliases: []string{"d"},
			Value:   MdDir,
			Usage:   "Markdown files dir",
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
	}

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
	}

	return &web
}
