package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type params struct {
	stackName   string
	domainName  string
	admunUser   string
	adminPw     string
	basicAuthPw string
	slackURL    string
	slackUser   string
	traefikPort string
	schema      string
	metricPort  string
	gwBridge    string
}

var questions = []string{
	"Enter docker stack name (default is swarmon): ",
	"Enter domain name: ",
	"Enter admin user name: ",
	"Enter admin password: ",
	"Enter basicAuth password (it will be hashed for Traefik):",
	"Enter Slack Webhook URL: ",
	"Enter username for Slack alerts: ",
	"Enter Traefik external port: ",
	"Enter HTTP schema (http or https): ",
	"Enter Docker Swarm metric port (you sould enable it manually): ",
	"Enter GW_BRIDGE ip (default is 172.18.0.1): ",
}

var length = len(questions)

func main() {
	answers := getAnswers()
	execAll(answers)
}

func execAll(answers []string) {
	fmt.Print("\nOutput:\n")
	for i := 0; i < length; i++ {
		if i == 0 {
			// printContainerID(answers[0])
		} else if i == 1 {
			// TODO
		}
	}
}

func getAnswers() []string {
	answers := make([]string, length)

	for i, question := range questions {
		answers[i] = readInput(question)
	}
	return answers
}

func readInput(question string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(question)

	var answer string
	if scanner.Scan() {
		answer = scanner.Text()
	}
	return answer
}

// func printContainerID(imageName string) {
// 	cli, err := client.NewEnvClient()
// 	try(err)

// 	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
// 	try(err)

// 	for _, container := range containers {
// 		if strings.Contains(container.Image, imageName) {
// 			fmt.Printf("%s\n", container.ID[:12])
// 		}
// 	}
// }

func deployStack() {
	exec.Command("docker", "stack deploy -c ./templates/compose.yml swarmon")
}

func try(err error) {
	if err != nil {
		panic(err)
	}
}
