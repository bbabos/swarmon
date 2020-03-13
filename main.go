package main

import (
	"fmt"
)

type input struct {
	Question string
	Answer   string
}

type param struct {
	Node struct {
		ID string
	}
	Tag                  string
	Domain               string
	AdminUser            string
	AdminPassword        string
	StackName            string
	Schema               string
	SlackURL             string
	SlackUser            string
	Port                 string
	MetricPort           string
	GwBridgeIP           string
	TraefikAdminPassword string
}

var p = param{Tag: "development", Node: struct{ ID string }{"{{.Node.ID}}"}}

var inputs = []input{
	{Question: "Docker stack name", Answer: "swarmon"},
	{Question: "Domain name", Answer: "localhost"},
	{Question: "Admin username", Answer: "admin"},
	{Question: "Admin password", Answer: "admin"},
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
	// gitClone("https://github.com/github/platform-samples.git")

	getAnswers()
	// printAll()

	result := parseFile("templates/example.yml", p)
	writeToFile(result, "templates/parsed.yml")
}

func notOkSolution() {
	p.StackName = inputs[0].Answer
	p.Domain = inputs[1].Answer
	p.AdminUser = inputs[2].Answer
	p.AdminPassword = inputs[3].Answer
	p.TraefikAdminPassword = inputs[4].Answer
	p.SlackURL = inputs[5].Answer
	p.SlackUser = inputs[6].Answer
	p.Port = inputs[7].Answer
	p.Schema = inputs[8].Answer
	p.MetricPort = inputs[9].Answer
	p.GwBridgeIP = inputs[10].Answer
}

func getAnswers() {
	for i := 0; i < length; i++ { // TODO i < length
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
	notOkSolution()
}

func printAll() {
	fmt.Print("\nAnswers:\n")
	for i := 0; i < length; i++ {
		if inputs[i].Answer != "" {
			fmt.Println(inputs[i].Answer)
		}
	}
}
