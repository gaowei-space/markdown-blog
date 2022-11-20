package main

import (
	"log"
	"os"

	"github.com/gaowei-space/markdown-blog/internal/app"
	"github.com/gaowei-space/markdown-blog/internal/utils"
	"github.com/urfave/cli"
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
	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getCommands() []cli.Command {
	command := cli.Command{
		Name:   "web",
		Usage:  "Run blog web server",
		Action: app.RunWeb,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "dir,d",
				Value: MdDir,
				Usage: "Markdown files dir",
			},
			cli.StringFlag{
				Name:  "title,t",
				Value: Title,
				Usage: "Blog title",
			},
			cli.IntFlag{
				Name:  "port,p",
				Value: DefaultPort,
				Usage: "Bind port",
			},
			cli.StringFlag{
				Name:  "env,e",
				Value: "prod",
				Usage: "Runtime environment, dev|test|prod",
			},
			cli.StringFlag{
				Name:  "index,i",
				Value: "",
				Usage: "Home page, default is empty",
			},
			cli.IntFlag{
				Name:  "cache,c",
				Value: 3,
				Usage: "The cache time unit is minutes, this parameter takes effect in the prod environment",
			},
			cli.StringFlag{
				Name:  "analyzer-baidu,ab",
				Value: "",
				Usage: "Set up Baidu Analyzer, default is empty",
			},
			cli.StringFlag{
				Name:  "analyzer-google,ag",
				Value: "",
				Usage: "Set up Google Analyzer, default is empty",
			},
		},
	}

	return []cli.Command{command}
}
