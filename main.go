package main

import (
	"fmt"
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
	{Question: "Docker gwbridge IP", Answer: "172.18.0.1"},
}
var length = len(inputs)
var p = param{Tag: "development", Node: struct{ ID string }{"{{.Node.ID}}"}}

func main() {
	menuPage()
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

func stackInit() {
	fmt.Println()
	fmt.Println("Swarm stack initialization started...")
	// gitClone("https://github.com/babobene/swarmon.git", "tmp")
	getAnswers()
	parsedfile := parseFile("templates/example.conf", p)
	writeToFile(parsedfile, "templates/parsed.yml")
	// stackDeploy("templates/parsed.yml", p.Docker.StackName)
}

func getAnswers() {
	for i := 0; i < length; i++ {
		if inputs[i].Answer == "" {
			inputs[i].Question = inputs[i].Question + ": "
			fmt.Print(inputs[i].Question)
			inputs[i].Answer = readInput()
		} else {
			inputs[i].Question = inputs[i].Question + " [" + inputs[i].Answer + "]" + ": "
			fmt.Print(inputs[i].Question)
			result := readInput()
			if result != "" {
				inputs[i].Answer = result
			}
		}
	}
	setParams()
}
