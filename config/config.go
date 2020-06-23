package config

import (
	"encoding/json"
	"io/ioutil"
)

type input struct {
	Question string
	Answer   string
}

// Param is ...
type Param struct {
	Tag    string
	Domain string
	Schema string
	Cgroup string
	Node   struct {
		ID string
	}
	AdminUser struct {
		Name     string
		Password string
	}
	Slack struct {
		Webhook   string
		AlertUser string
	}
	Traefik struct {
		Port       string
		BAPassword string
		BAUser     string
	}
	Docker struct {
		StackName  string
		MetricPort string
		GwBridgeIP string
	}
	HostNamePath string // for dev only
}

// Inputs is ...
var Inputs = []input{
	{Question: "Docker stack name"},
	{Question: "Domain name"},
	{Question: "Admin username"},
	{Question: "Admin password"},
	{Question: "BasicAuth username"},
	{Question: "BasicAuth password"},
	{Question: "Slack Webhook URL"},
	{Question: "Username for Slack alerts"},
	{Question: "Traefik external port"},
	{Question: "HTTP schema"},
	{Question: "Docker Swarm metric port"},
	{Question: "Docker gwbridge IP"},
}

// Params is ...
var Params = Param{
	Tag:    "development",
	Node:   struct{ ID string }{"{{.Node.ID}}"},
	Cgroup: "# - /cgroup:/sys/fs/cgroup:ro",
}

// Path is ...
var Path = "config/stackconfig.json"

// Save is ...
func Save(folderPath string) {
	file, _ := json.MarshalIndent(Params, "", " ")
	_ = ioutil.WriteFile(folderPath, file, 0644)
}

// Load is ...
func Load(filePath string) {
	file, _ := ioutil.ReadFile(filePath)
	_ = json.Unmarshal([]byte(file), &Params)
}
