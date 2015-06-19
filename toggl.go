package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/mozillazg/request"
	"github.com/spf13/viper"
)

// Toggl base url
const (
	BaseURL       = "https://toggl.com/reports/api/v2/details"
	togglTimeForm = "2006-01-02T15:04:05-07:00"
)

// Toggl ...It is main function
func Toggl() {

	// name of config file (without extension)
	viper.SetConfigName("config")
	// call multiple times to add many search paths
	viper.AddConfigPath("$GOPATH/src/github.com/sotayamashita/report")
	// Find and read the config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 2. fetch data based on config.toml
	client := new(http.Client)
	req := request.NewRequest(client)
	req.BasicAuth = request.BasicAuth{viper.GetString("toggl.api_token"), "api_token"}
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
	t, _ := time.Parse(togglTimeForm, date.(string))
	return t.Format("3:4 PM")
}

// TODO: methodize like URL
func urlBuilder() *url.URL {

	u, err := url.Parse(BaseURL)
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Add("user_agent", "report")
	q.Add("workspace_id", viper.GetString("toggl.workspace_id"))
	q.Add("since", time.Now().Format("2006-01-02"))
	u.RawQuery = q.Encode()

	return u
}
