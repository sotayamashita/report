package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "report"
	app.Version = Version
	app.Usage = ""
	app.Author = "Sota Yamashita"
	app.Email = "sota.yamashita@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
