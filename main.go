package main

import (
	"log"
	"os"

	"github.com/gaowei-space/markdown-blog/internal/app"
	"github.com/gaowei-space/markdown-blog/internal/utils"
	"github.com/urfave/cli"
)

var (
	MdDir                = "md/"
	Title                = "Blog"
	AppVersion           = "0.0.3"
	BuildDate, GitCommit string
)

// web服务器默认端口
const DefaultPort = 5006

type navItem struct {
	Name  string
	Path  string
	IsDir bool
}

func main() {
	cliApp := cli.NewApp()
	cliApp.Name = "markdown-blog"
	cliApp.Usage = "Markdown Blog Service"
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
				Usage: "markdown files dir",
			},
			cli.StringFlag{
				Name:  "title,t",
				Value: Title,
				Usage: "blog title, default is Blog",
			},
			cli.IntFlag{
				Name:  "port,p",
				Value: DefaultPort,
				Usage: "bind port",
			},
			cli.StringFlag{
				Name:  "env,e",
				Value: "prod",
				Usage: "runtime environment, dev|test|prod",
			},
			cli.StringFlag{
				Name:  "index,i",
				Value: "",
				Usage: "home page, default is empty",
			},
			cli.StringFlag{
				Name:  "analyzer-baidu",
				Value: "",
				Usage: "Set up Baidu Analyzer, default is empty",
			},
			cli.StringFlag{
				Name:  "analyzer-google",
				Value: "",
				Usage: "Set up Google Analyzer, default is empty",
			},
		},
	}

	return []cli.Command{command}
}
