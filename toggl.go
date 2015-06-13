package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/mozillazg/request"
)

// Toggl base url
const BaseURL = "https://toggl.com/reports/api/v2/details"
const togglForm = "2006-01-02T15:04:05-07:00"

// TogglConfig ...
type TogglConfig struct {
	APIToken    string `toml:"api_token"`
	WorkspaceID string `toml:"workspace_id"`
}

// Toggl ...It is main function
func Toggl() {
	var config Config

	// 1. read config.toml
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
		log.Fatal(err.Error())
		return
	}
	j, err := resp.Json()
	defer resp.Body.Close()

	// 3. process data and return
	// TODO: methodize processror
	// tasks
	tasks, _ := j.Get("data").Array()
	for n := range tasks {
		show(tasks[n])
	}

}

func show(task interface{}) {
	// Assert task intarface{} to map[string] interface{}
	var t = task.(map[string]interface{})

	var tasks []interface{}
	tasks = append(tasks, t["description"])
	tasks = append(tasks, t["tags"])
	tasks = append(tasks, t["project"])
	tasks = append(tasks, t["client"])
	tasks = append(tasks, duration(t["dur"]))
	tasks = append(tasks, meridiemTime(t["start"]))
	tasks = append(tasks, "-")
	tasks = append(tasks, meridiemTime(t["end"]))

	fmt.Println(tasks)
}

func duration(duration interface{}) time.Duration {
	// Assert duration interface{} to json.Number
	var o, _ = duration.(json.Number)
	// Assert originalDuration json.Number to int64
	var i, _ = o.Int64()
	// Return time.Duration
	return time.Duration(i) * time.Millisecond
}

func meridiemTime(date interface{}) string {
	t, _ := time.Parse(togglForm, date.(string))
	return t.Format("3:4 PM")
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
	q.Add("since", time.Now().Format("2006-01-02"))
	u.RawQuery = q.Encode()

	return u
}
