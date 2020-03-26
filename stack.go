package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

var configPath = "config.json"
var parsedStackFilePath = "tmp/parsed.yml"
var rawStackFilePath = "tmp/docker-compose.yml"

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
	p.Traefik.BAPassword = hashPass(inputs[5].Answer) // TODO
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
	clear()
	fmt.Println("-------------------------------------")
	fmt.Println("Swarm stack initialization started...")
	fmt.Println("-------------------------------------")

	gitClone("https://github.com/babobene/swarmon.git", "tmp")
	getAnswers()
	parsedFile := parseFile(rawStackFilePath, p)
	writeToFile(parsedFile, parsedStackFilePath)

	if stackExist() {
		fmt.Println("\nUpdating monitoring services...")
		fmt.Println("-------------------------------")
	} else {
		fmt.Println("\nStack deploy started...")
		fmt.Println("-----------------------")
	}
	execCommand("docker stack deploy -c " + parsedStackFilePath + " " + p.Docker.StackName)
}

func stackDelete() {
	if stackExist() {
		execCommand("docker stack rm " + p.Docker.StackName)
	} else {
		fmt.Println("You may not have a monitoring stack deployed!")
	}
}

func stackExist() bool {
	cmd := exec.Command("docker", "stack", "ls", "--format", "'{{.Name}}'")
	var out bytes.Buffer
	cmd.Stdout = &out

	cmd.Run()
	stdout := out.String()

	if strings.Contains(stdout, p.Docker.StackName) {
		return true
	}
	return false
}
