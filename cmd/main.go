package main

import (
	"github.com/clean_project_frame/commands"
	"github.com/urfave/cli"
	"os"
)

var version = "development"

func main() {
	app := cli.NewApp()
	app.Name = "clean_project_frame"
	app.Usage = ""
	app.Version = version
	app.Copyright = "(c) 2019 The clean_project_frame contributors <fionawp@126.com>"
	app.EnableBashCompletion = true
	app.Flags = commands.GlobalFlags

	app.Commands = []cli.Command{
		commands.ConfigCommand,
		commands.StartCommand,
	}

	app.Run(os.Args)
}
