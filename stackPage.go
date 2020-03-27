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
	length := len(inputs)
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
	var selected string
	stackexist := stackExist()
	exit := false

	clear()
	if stackexist {
		fmt.Println("-----------------------------------")
		fmt.Println("Update existing monitoring stack...")
		fmt.Println("-----------------------------------")
	} else {
		fmt.Println("----------------------------------------------")
		fmt.Println("New monitoring stack initialization started...")
		fmt.Println("----------------------------------------------")
	}

	gitClone("https://github.com/babobene/swarmon.git", "tmp")
	getAnswers()
	parsedFile := parseFile(rawStackFilePath, p)
	writeToFile(parsedFile, parsedStackFilePath)

	if stackexist {
		fmt.Println("\nUpdating docker services...")
		fmt.Println("-------------------------------")
	} else {
		fmt.Println("\nStack deploy started...")
		fmt.Println("-----------------------")
	}
	execCommand("docker stack deploy -c " + parsedStackFilePath + " " + p.Docker.StackName)

	for !exit {
		fmt.Print("Enter 0 to exit: ")
		selected = readInput()

		switch selected {
		case "0":
			exit = true
			menuPage()
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}

func stackDelete() {
	stackexist := stackExist()

	if stackexist {
		fmt.Print("Are you sure? [y/N]: ")
		input := readInput()
		if input == "y" {
			execCommand("docker stack rm " + p.Docker.StackName)
			fmt.Println("Monitoring stack deleted.")
		}
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

	contains := strings.Contains(stdout, p.Docker.StackName)
	if contains {
		return true
	}
	return false
}

func checkPreviousMonStack() bool {
	configexist := fileExists(configPath)

	if configexist {
		stackexist := stackExist()

		if stackexist {
			fmt.Printf("You have a previously deployed monitoring stack (%s)!\n", p.Docker.StackName)
			return true
		}
	}
	return false
}
