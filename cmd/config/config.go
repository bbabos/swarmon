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
	StackConfig: "internal/stackconfig.json",
	RawStack:    "internal/docker/docker-compose.yml",
	ParsedStack: "internal/docker/parsed.yml",
}

// Save is ...
func Save() {
	data, _ := json.MarshalIndent(Params, "", " ")
	_ = ioutil.WriteFile(Paths.StackConfig, data, 0644)
}

// Load is ...
func Load(filePath string) {
	file, _ := ioutil.ReadFile(filePath)
	_ = json.Unmarshal([]byte(file), &Params)
}

// SetAnswers is ...
func SetAnswers() {
	Inputs[0].Answer = Params.Docker.StackName
	Inputs[1].Answer = Params.Domain
	Inputs[2].Answer = Params.AdminUser.Name
	Inputs[3].Answer = Params.AdminUser.Password
	Inputs[4].Answer = Params.Traefik.BAUser
	Inputs[5].Answer = Params.Traefik.BAPassword
	Inputs[6].Answer = Params.Slack.Webhook
	Inputs[7].Answer = Params.Slack.AlertUser
	Inputs[8].Answer = Params.Traefik.Port
	Inputs[9].Answer = Params.Schema
	Inputs[10].Answer = Params.Docker.MetricPort
	Inputs[11].Answer = Params.Docker.GwBridgeIP
}

// SetParams is ...
func SetParams() {
	Params.Docker.StackName = Inputs[0].Answer
	Params.Domain = Inputs[1].Answer
	Params.AdminUser.Name = Inputs[2].Answer
	Params.AdminUser.Password = Inputs[3].Answer
	Params.Traefik.BAUser = Inputs[4].Answer
	Params.Traefik.BAPassword = Inputs[5].Answer
	Params.Slack.Webhook = Inputs[6].Answer
	Params.Slack.AlertUser = Inputs[7].Answer
	Params.Traefik.Port = Inputs[8].Answer
	Params.Schema = Inputs[9].Answer
	Params.Docker.MetricPort = Inputs[10].Answer
	Params.Docker.GwBridgeIP = Inputs[11].Answer
}
