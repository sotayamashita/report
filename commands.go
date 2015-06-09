package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/codegangsta/cli"
	"github.com/mozillazg/request"
)

var Commands = []cli.Command{
	commandToggle,
}

var commandToggle = cli.Command{
	Name:  "toggl",
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
	Toggl Toggl
}

type Toggl struct {
	APIToken string `toml:"api_token"`
}

func doToggle(c *cli.Context) {
	var config Config
	_, err := toml.DecodeFile("config.tml", &config)
	if err != nil {
		panic(err)
	}

	// debug
	fmt.Printf("%s\n", config.Toggl.APIToken)

	// Get toggl api
	client := new(http.Client)
	req := request.NewRequest(client)
	req.BasicAuth = request.BasicAuth{config.Toggl.APIToken, "api_token"}
	resp, err := req.Get("https://toggl.com/reports/api/v2")
	j, err := resp.Json()
	defer resp.Body.Close()

	fmt.Println(j)

}
