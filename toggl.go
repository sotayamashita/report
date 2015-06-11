package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/BurntSushi/toml"
	"github.com/mozillazg/request"
)

// Toggl base url
const BaseURL = "https://toggl.com/reports/api/v2/details"

// TogglConfig ...
type TogglConfig struct {
	APIToken    string `toml:"api_token"`
	WorkspaceID string `toml:"workspace_id"`
}

// Toggl ...It is main function
func Toggl() {

	// 1. read config.toml
	var config Config
	_, err := toml.DecodeFile("config.tml", &config)
	if err != nil {
		panic(err)
	}

	// 2. fetch data based on config.toml
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

	// 3. process data and return
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
