package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/bbabos/swarmon/cmd/utils"
)

// Inputs is ...
var Inputs = []input{
	{Question: "Docker stack name"},
	{Question: "SwarMon Docker tag"},
	{Question: "Domain name"},
	{Question: "Prometheus subdomain"},
	{Question: "Grafana subdomain"},
	{Question: "Alertmanager subdomain"},
	{Question: "Admin user for Grafana"},
	{Question: "Admin password for Grafana"},
	{Question: "BasicAuth username"},
	{Question: "BasicAuth password"},
	{Question: "Slack Webhook URL"},
	{Question: "Username for Slack alerts"},
	{Question: "Channel for Slack alerts"},
	{Question: "Traefik external port"},
	{Question: "HTTP schema"},
	{Question: "Docker Swarm metric port"},
	{Question: "Docker gwbridge IP"},
	{Question: "Enable/disable cgroup (y/n)"},
	{Question: "Cgroup path"},
	{Question: "Hostname path"},
}

// Params is ...
var Params = params{
	Node: struct{ ID string }{"{{.Node.ID}}"},
	Docker: struct {
		Tag        string
		StackName  string
		MetricPort string
		GwBridgeIP string
	}{"master", "", "9323", "172.18.0.1"},
	Cgroup: struct {
		Path    string
		Enabled string
	}{"/cgroup", "n"},
	HostNamePath: "/etc/hostname",
}

// Paths is ...
var Paths = paths{
	StackConfig: "stackconfig.json",
	RawStack:    "configs/docker/docker-compose.yml",
	ParsedStack: "configs/docker/parsed.yml",
}

// CreateOrSave is ...
func CreateOrSave(configPath string) {
	data, _ := json.MarshalIndent(Params, "", " ")
	_ = ioutil.WriteFile(configPath, data, 0644)
}

// Load is ...
func Load(filePath string) {
	file, _ := ioutil.ReadFile(filePath)
	_ = json.Unmarshal([]byte(file), &Params)
}

// GetAnswers is ...
func GetAnswers(stackExists bool) {
	length := len(Inputs)
	num := 0
	if stackExists {
		num = 1
	}
	for i := num; i < length; i++ {
		var question string
		if Inputs[i].Answer == "" {
			question = Inputs[i].Question + ": "
			fmt.Print(question)
			Inputs[i].Answer = utils.ReadInput()
		} else {
			question = Inputs[i].Question + " [" + Inputs[i].Answer + "]" + ": "
			fmt.Print(question)
			result := utils.ReadInput()
			if result != "" {
				Inputs[i].Answer = result
			}
		}
	}
	SetParams()
	CreateOrSave(Paths.StackConfig)
	Params.Traefik.BAPassword = utils.HashPass(Inputs[9].Answer)
	if Inputs[17].Answer == "y" {
		Params.Cgroup.Enabled = "-"
	} else if Inputs[17].Answer == "n" {
		Params.Cgroup.Enabled = "#-"
	}
}

// SetAnswers is ...
func SetAnswers() {
	Inputs[0].Answer = Params.Docker.StackName
	Inputs[1].Answer = Params.Docker.Tag
	Inputs[2].Answer = Params.Domain
	Inputs[3].Answer = Params.Traefik.PrometheusSubDomain
	Inputs[4].Answer = Params.Traefik.GrafanaSubDomain
	Inputs[5].Answer = Params.Traefik.AlertmanagerSubDomain
	Inputs[6].Answer = Params.AdminUser.Name
	Inputs[7].Answer = Params.AdminUser.Password
	Inputs[8].Answer = Params.Traefik.BAUser
	Inputs[9].Answer = Params.Traefik.BAPassword
	Inputs[10].Answer = Params.Slack.Webhook
	Inputs[11].Answer = Params.Slack.AlertUser
	Inputs[12].Answer = Params.Slack.Channel
	Inputs[13].Answer = Params.Traefik.Port
	Inputs[14].Answer = Params.Schema
	Inputs[15].Answer = Params.Docker.MetricPort
	Inputs[16].Answer = Params.Docker.GwBridgeIP
	Inputs[17].Answer = Params.Cgroup.Enabled
	Inputs[18].Answer = Params.Cgroup.Path
	Inputs[19].Answer = Params.HostNamePath
}

// SetParams is ...
func SetParams() {
	Params.Docker.StackName = Inputs[0].Answer
	Params.Docker.Tag = Inputs[1].Answer
	Params.Domain = Inputs[2].Answer
	Params.Traefik.PrometheusSubDomain = Inputs[3].Answer
	Params.Traefik.GrafanaSubDomain = Inputs[4].Answer
	Params.Traefik.AlertmanagerSubDomain = Inputs[5].Answer
	Params.AdminUser.Name = Inputs[6].Answer
	Params.AdminUser.Password = Inputs[7].Answer
	Params.Traefik.BAUser = Inputs[8].Answer
	Params.Traefik.BAPassword = Inputs[9].Answer
	Params.Slack.Webhook = Inputs[10].Answer
	Params.Slack.AlertUser = Inputs[11].Answer
	Params.Slack.Channel = Inputs[12].Answer
	Params.Traefik.Port = Inputs[13].Answer
	Params.Schema = Inputs[14].Answer
	Params.Docker.MetricPort = Inputs[15].Answer
	Params.Docker.GwBridgeIP = Inputs[16].Answer
	Params.Cgroup.Enabled = Inputs[17].Answer
	Params.Cgroup.Path = Inputs[18].Answer
	Params.HostNamePath = Inputs[19].Answer
}
