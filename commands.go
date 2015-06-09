package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandToggle,
}

var commandToggle = cli.Command{
	Name:  "toggle",
	Usage: "",
	Description: `
`,
	Action: doToggle,
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

type Config struct {
	Toggle Toggle
}

type Toggle struct {
	ApiToken string `toml:"api_token"`
}

func doToggle(c *cli.Context) {
	var config Config
	_, err := toml.DecodeFile("config.tml", &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", config.Toggle.ApiToken)
}
