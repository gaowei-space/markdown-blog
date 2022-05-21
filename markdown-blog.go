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
	AppVersion           = "0.0.1"
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
	cliApp.Name = "GoBlog"
	cliApp.Usage = "GoBlog Service"
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
		Usage:  "run web server",
		Action: app.RunWeb,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "dir,d",
				Value: MdDir,
				Usage: "markdown files dir",
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
		},
	}

	return []cli.Command{command}
}
