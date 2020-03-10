package main

import (
	"bufio"
	"fmt"
	"os"
)

type input struct {
	Question string
	Answer   string
}

func (i input) execute(action func(text string) string) {
	action(i.Answer)
}

var inputs = []input{
	{Question: "Docker stack name", Answer: "random"},
	{Question: "Domain name"},
	{Question: "Admin username"},
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
	getAnswers()
	executeAll()
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

func readInput(question string) string {
	scanner := bufio.NewScanner(os.Stdin)
	var answer string

	if scanner.Scan() {
		answer = scanner.Text()
	}
	return answer
}

func executeAll() {
	fmt.Print("\nOutput:\n")
	for i := 0; i < length; i++ {
		if inputs[i].Answer != "" {
			fmt.Println(inputs[i].Answer)
		}
	}
}

func try(err error) {
	if err != nil {
		panic(err)
	}
}
