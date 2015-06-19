package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

// Commands ...
var Commands = []cli.Command{
	commandToggle,
}

var commandToggle = cli.Command{
	Name:        "toggl",
	Usage:       "Shows your activity with Toggl",
	Description: `Shows your activity with Toggl on today`,
	Action:      doToggle,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doToggle(c *cli.Context) {
	Toggl()
}
