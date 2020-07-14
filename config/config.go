package config

import (
	"encoding/json"
	"io/ioutil"
)

type input struct {
	Question string
	Answer   string
}

type params struct {
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
	HostNamePath string // dev only
}

type paths struct {
	StackConfig string
	RawStack    string
	ParsedStack string
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
var Params = params{
	Tag:    "development",
	Node:   struct{ ID string }{"{{.Node.ID}}"},
	Cgroup: "# - /cgroup:/sys/fs/cgroup:ro",
}

// Paths is ...
var Paths = paths{
	StackConfig: "config/stackconfig.json",
	RawStack:    "config/docker/docker-compose.yml",
	ParsedStack: "config/docker/parsed.yml",
}

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
