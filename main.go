package main

import (
	"fmt"

	"github.com/docker/docker/client"
)

type input struct {
	Question string
	Answer   string
}

type param struct {
	Tag    string
	Domain string
	Schema string
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
}

var p = param{Tag: "development", Node: struct{ ID string }{"{{.Node.ID}}"}}

var inputs = []input{
	{Question: "Docker stack name", Answer: "swarmon"},
	{Question: "Domain name", Answer: "localhost"},
	{Question: "Admin username", Answer: "admin"},
	{Question: "Admin password", Answer: "adminpw"},
	{Question: "BasicAuth username", Answer: "admin"},
	{Question: "BasicAuth password", Answer: "hashedpw"},
	{Question: "Slack Webhook URL", Answer: "http://webhook.url.com"},
	{Question: "Username for Slack alerts", Answer: "alertmanager"},
	{Question: "Traefik external port", Answer: "80"},
	{Question: "HTTP schema", Answer: "http"},
	{Question: "Docker Swarm metric port", Answer: "9323"},
	{Question: "GW_BRIDGE IP", Answer: "172.18.0.1"},
}
var length = len(inputs)

func main() {
	cli, err := client.NewEnvClient()
	try(err)

	listServices(cli)
	listSwarmNodes(cli)
	// getAnswers()
	// parsedfile := parseFile("templates/example.yml", p)
	// writeToFile(parsedfile, "templates/parsed.yml")
}

func setParams() {
	p.Docker.StackName = inputs[0].Answer
	p.Domain = inputs[1].Answer
	p.AdminUser.Name = inputs[2].Answer
	p.AdminUser.Password = inputs[3].Answer
	p.Traefik.BAUser = inputs[4].Answer
	p.Traefik.BAPassword = hashPass(inputs[5].Answer)
	p.Slack.Webhook = inputs[6].Answer
	p.Slack.AlertUser = inputs[7].Answer
	p.Traefik.Port = inputs[8].Answer
	p.Schema = inputs[9].Answer
	p.Docker.MetricPort = inputs[10].Answer
	p.Docker.GwBridgeIP = inputs[11].Answer
}

func getAnswers() {
	for i := 0; i < length; i++ {
		if inputs[i].Answer == "" {
			inputs[i].Question = inputs[i].Question + ": "
			fmt.Print(inputs[i].Question)
			inputs[i].Answer = readInput(inputs[i].Question)
		} else {
			inputs[i].Question = inputs[i].Question + " [" + inputs[i].Answer + "]" + ": "
			fmt.Print(inputs[i].Question)
			result := readInput(inputs[i].Question)
			if result != "" {
				inputs[i].Answer = result
			}
		}
	}
	setParams()
}
