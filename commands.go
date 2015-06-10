package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/codegangsta/cli"
	"github.com/mozillazg/request"
)

// Commands ...
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

// Config ...
type Config struct {
	Toggl TogglConfig
}

// TogglConfig ...
type TogglConfig struct {
	APIToken    string `toml:"api_token"`
	WorkspaceID string `toml:"workspace_id"`
}

func doToggle(c *cli.Context) {
	// TODO: standarize or make it like global
	// TODO: methodize like import
	var config Config
	_, err := toml.DecodeFile("config.tml", &config)
	if err != nil {
		panic(err)
	}

	// TODO: methodize like fetch
	client := new(http.Client)
	req := request.NewRequest(client)
	req.BasicAuth = request.BasicAuth{config.Toggl.APIToken, "api_token"}
	resp, err := req.Get(urlBuilder())
	if err != nil {
		log.Fatal(err)
		return
	}
	j, err := resp.Json()
	defer resp.Body.Close() // Do not forget close body

	// TODO: methodize processror
	// tasks
	tasks, _ := j.Get("data").Array()

	// Init new tasks
	// TODO: init new with key
	var newTasks = make(map[string]interface{})
	newTasks["description"] = ""
	newTasks["clinet"] = ""
	newTasks["project"] = ""
	newTasks["duration"] = ""

	// print each task
	for n := range tasks {
		castedTask, _ := tasks[n].(map[string]interface{})
		fmt.Println(castedTask["description"])
	}

}

// Toggl base url
const BaseURL = "https://toggl.com/reports/api/v2/details"

// TODO: methodize like URL
func urlBuilder() *url.URL {
	var config Config
	_, err := toml.DecodeFile("config.tml", &config)
	if err != nil {
		panic(err)
	}

	u, err := url.Parse(BaseURL)
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Add("user_agent", "report")
	q.Add("workspace_id", config.Toggl.WorkspaceID)
	u.RawQuery = q.Encode()

	return u
}
