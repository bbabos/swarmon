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
var length = len(inputs)
var p = param{Tag: "development", Node: struct{ ID string }{"{{.Node.ID}}"}}
var configPath = "templates/config.json"

func main() {
	if fileExists(configPath) {
		loadConfig(configPath)
		setAnswers()
	}
	menuPage()
}

func getAnswers() {
	for i := 0; i < length; i++ {
		if inputs[i].Answer == "" {
			fmt.Print(inputs[i].Question + ": ")
			inputs[i].Answer = readInput()
		} else {
			fmt.Print(inputs[i].Question + " [" + inputs[i].Answer + "]" + ": ")
			result := readInput()
			if result != "" {
				inputs[i].Answer = result
			}
		}
	}
	setParams()
	saveConfig(configPath)
}

func setAnswers() {
	inputs[0].Answer = p.Docker.StackName
	inputs[1].Answer = p.Domain
	inputs[2].Answer = p.AdminUser.Name
	inputs[3].Answer = p.AdminUser.Password
	inputs[4].Answer = p.Traefik.BAUser
	inputs[5].Answer = p.Traefik.BAPassword
	inputs[6].Answer = p.Slack.Webhook
	inputs[7].Answer = p.Slack.AlertUser
	inputs[8].Answer = p.Traefik.Port
	inputs[9].Answer = p.Schema
	inputs[10].Answer = p.Docker.MetricPort
	inputs[11].Answer = p.Docker.GwBridgeIP
}

func setParams() {
	p.Docker.StackName = inputs[0].Answer
	p.Domain = inputs[1].Answer
	p.AdminUser.Name = inputs[2].Answer
	p.AdminUser.Password = inputs[3].Answer
	p.Traefik.BAUser = inputs[4].Answer
	p.Traefik.BAPassword = inputs[5].Answer
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
	gitClone("https://github.com/babobene/swarmon.git", "tmp")
	getAnswers()
	p.Traefik.BAPassword = hashPass(inputs[5].Answer) // TODO
	parsedfile := parseFile("tmp/docker-compose.yml", p)
	writeToFile(parsedfile, "tmp/parsed.yml")
	fmt.Println()
	deployStack()
}
