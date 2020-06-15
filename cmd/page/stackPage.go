package page

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/bbabos/swarmon/cmd/utils"
	"github.com/bbabos/swarmon/config"
)

var rawStackFilePath = "config/docker/docker-compose.yml"
var parsedStackFilePath = "config/docker/parsed.yml"
var stackexist = stackExist()

func getAnswers(stackExists bool) {
	length := len(config.Inputs)
	var num int
	if stackExists {
		num = 1
	} else {
		num = 0
	}
	for i := num; i < length; i++ {
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
	config.Save(config.Path)
	config.Params.Traefik.BAPassword = utils.HashPass(config.Inputs[5].Answer)
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

func stackInitOrUpdate() {
	stackexist = stackExist()
	if stackexist {
		fmt.Println("-----------------------------------")
		fmt.Println("Update existing monitoring stack...")
		fmt.Println("-----------------------------------")
	} else {
		fmt.Println("----------------------------------------------")
		fmt.Println("New monitoring stack initialization started...")
		fmt.Println("----------------------------------------------")
	}

	getAnswers(stackexist)
	parsedFile := utils.ParseFile(rawStackFilePath, config.Params)
	utils.WriteToFile(parsedFile, parsedStackFilePath)

	if stackexist {
		fmt.Println("---------------------------")
		fmt.Println("Updating docker services...")
		fmt.Println("---------------------------")
	} else {
		fmt.Println("-----------------------")
		fmt.Println("Stack deploy started...")
		fmt.Println("-----------------------")
	}

	utils.ExecShellCommand("docker stack deploy -c " + parsedStackFilePath + " " + config.Params.Docker.StackName)
	utils.ExitOnKeyStroke(stackPage)
}

func stackDelete() {
	stackexist = stackExist()
	if stackexist {
		fmt.Print("Are you sure? [y/N]: ")
		input := utils.ReadInput()
		if input == "y" {
			utils.ExecShellCommand("docker stack rm " + config.Params.Docker.StackName)
			fmt.Println("-------------------------------")
			fmt.Println("Monitoring stack deleted successfully!")
		}
	} else {
		fmt.Println("You may not have a monitoring stack deployed!")
	}
	utils.ExitOnKeyStroke(stackPage)
}

func stackExist() bool {
	cmd := exec.Command("docker", "stack", "ls", "--format", "'{{.Name}}'")
	var out bytes.Buffer
	cmd.Stdout = &out

	cmd.Run()
	stdout := out.String()

	contains := strings.Contains(stdout, config.Params.Docker.StackName)
	return contains
}
