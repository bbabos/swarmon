package main

import (
	"fmt"
)

type input struct {
	Question string
	Answer   string
	Param    param
}

type param struct {
	Tag                  string
	Domain               string
	AdminUser            string
	AdminPassword        string
	StackName            string
	Schema               string
	SlackURL             string
	SlackUser            string
	Port                 int
	MetricPort           int
	GwBridgeIP           string
	TraefikAdminPassword string
}

func (i *input) execute(action func(text string) string) {
	action(i.Answer)
}

var inputs = []input{
	{Question: "Docker stack name", Answer: "stackname", Param: param{StackName: ""}}, // TODO remove answer
	{Question: "Domain name", Answer: "domain name", Param: param{Domain: ""}},
	{Question: "Admin username", Param: param{AdminUser: ""}},
	{Question: "Admin password"},
	{Question: "BasicAuth password"},
	{Question: "Slack Webhook URL"},
	{Question: "Username for Slack alerts"},
	{Question: "Traefik external port"},
	{Question: "HTTP schema"},
	{Question: "Docker Swarm metric port"},
	{Question: "GW_BRIDGE IP"},
}
var length = len(inputs)

func main() {
	// gitClone("https://github.com/github/platform-samples.git")
	getAnswers()
	printAll()

	// testObj := input{Question: "Kerdes ehh", Answer: "Valasz ahh"}
	// result := parseFile("test-tmpl.yml", testObj)
	// writeToFile(result, "asd.txt")

	// fmt.Println(result)
}

func getAnswers() {
	for i := 0; i < 3; i++ { // TODO i < length
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
}

func printAll() {
	fmt.Print("\nAnswers:\n")
	for i := 0; i < length; i++ {
		if inputs[i].Answer != "" {
			fmt.Println(inputs[i].Answer)
		}
	}
}
