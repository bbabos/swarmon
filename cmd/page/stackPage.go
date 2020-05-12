package page

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/bbabos/swarmon/cmd/utils"
	"github.com/bbabos/swarmon/config"
)

var parsedStackFilePath = "tmp/parsed.yml"
var rawStackFilePath = "tmp/docker-compose.yml"

func stackPage() {
	var selected string
	items := []string{
		"1. Docker stack deploy/update",
		"2. Remove monitoring stack",
		"0. Back",
	}
	renderMenu(items, "STACK MENU")

loop:
	for {
		fmt.Print("Select an option: ")
		selected = utils.ReadInput()

		switch selected {
		case "1":
			StackInit()
			break loop
		case "2":
			StackDelete()
		case "0":
			MenuPage()
			break loop
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}

func getAnswers() {
	length := len(config.Inputs)
	for i := 0; i < length; i++ {
		if config.Inputs[i].Answer == "" {
			fmt.Print(config.Inputs[i].Question + ": ")
			config.Inputs[i].Answer = utils.ReadInput()
		} else {
			fmt.Print(config.Inputs[i].Question + " [" + config.Inputs[i].Answer + "]" + ": ")
			result := utils.ReadInput()
			if result != "" {
				config.Inputs[i].Answer = result
			}
		}
	}
	setParams()
	utils.SaveConfig(config.ConfigPath)
	config.Params.Traefik.BAPassword = utils.HashPass(config.Inputs[5].Answer) // TODO
}

// SetAnswers is ...
func SetAnswers() {
	config.Inputs[0].Answer = config.Params.Docker.StackName
	config.Inputs[1].Answer = config.Params.Domain
	config.Inputs[2].Answer = config.Params.AdminUser.Name
	config.Inputs[3].Answer = config.Params.AdminUser.Password
	config.Inputs[4].Answer = config.Params.Traefik.BAUser
	config.Inputs[5].Answer = config.Params.Traefik.BAPassword
	config.Inputs[6].Answer = config.Params.Slack.Webhook
	config.Inputs[7].Answer = config.Params.Slack.AlertUser
	config.Inputs[8].Answer = config.Params.Traefik.Port
	config.Inputs[9].Answer = config.Params.Schema
	config.Inputs[10].Answer = config.Params.Docker.MetricPort
	config.Inputs[11].Answer = config.Params.Docker.GwBridgeIP
}

func setParams() {
	config.Params.Docker.StackName = config.Inputs[0].Answer
	config.Params.Domain = config.Inputs[1].Answer
	config.Params.AdminUser.Name = config.Inputs[2].Answer
	config.Params.AdminUser.Password = config.Inputs[3].Answer
	config.Params.Traefik.BAUser = config.Inputs[4].Answer
	config.Params.Traefik.BAPassword = config.Inputs[5].Answer
	config.Params.Slack.Webhook = config.Inputs[6].Answer
	config.Params.Slack.AlertUser = config.Inputs[7].Answer
	config.Params.Traefik.Port = config.Inputs[8].Answer
	config.Params.Schema = config.Inputs[9].Answer
	config.Params.Docker.MetricPort = config.Inputs[10].Answer
	config.Params.Docker.GwBridgeIP = config.Inputs[11].Answer
}

// StackInit is ...
func StackInit() {
	var selected string
	stackexist := stackExist()

	utils.Clear()
	if stackexist {
		fmt.Println("-----------------------------------")
		fmt.Println("Update existing monitoring stack...")
		fmt.Println("-----------------------------------")
	} else {
		fmt.Println("----------------------------------------------")
		fmt.Println("New monitoring stack initialization started...")
		fmt.Println("----------------------------------------------")
	}

	utils.GitClone("https://github.com/babobene/swarmon.git", "tmp")
	getAnswers()
	parsedFile := utils.ParseFile(rawStackFilePath, config.Params)
	utils.WriteToFile(parsedFile, parsedStackFilePath)

	if stackexist {
		fmt.Println("\nUpdating docker services...")
		fmt.Println("-------------------------------")
	} else {
		fmt.Println("\nStack deploy started...")
		fmt.Println("-----------------------")
	}
	utils.ExecCommand("docker stack deploy -c " + parsedStackFilePath + " " + config.Params.Docker.StackName)

loop:
	for {
		fmt.Print("Enter 0 to exit: ")
		selected = utils.ReadInput()

		switch selected {
		case "0":
			MenuPage()
			break loop
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}

// StackDelete is ...
func StackDelete() {
	stackexist := stackExist()

	if stackexist {
		fmt.Print("Are you sure? [y/N]: ")
		input := utils.ReadInput()
		if input == "y" {
			utils.ExecCommand("docker stack rm " + config.Params.Docker.StackName)
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

	contains := strings.Contains(stdout, config.Params.Docker.StackName)
	if contains {
		return true
	}
	return false
}

func checkPreviousMonStack() bool {
	configexist := utils.FileExists(config.ConfigPath)

	if configexist {
		stackexist := stackExist()

		if stackexist {
			fmt.Printf("You have a previously deployed monitoring stack (%s)!\n", config.Params.Docker.StackName)
			return true
		}
	}
	return false
}
