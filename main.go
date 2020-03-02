package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var questions = []string{
	"Enter image name: ",
	"2nd question: ",
	"Third question: ",
}

var length = len(questions)

func main() {
	answers := getAnswers()
	execAll(answers)
}

func execAll(answers []string) {
	fmt.Print("\nConfig started...\n")
	for _, answer := range answers {
		printContainerID(answer)
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

func printContainerID(imageName string) {
	cli, err := client.NewEnvClient()
	try(err)

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	try(err)

	for _, container := range containers {
		if strings.Contains(container.Image, imageName) {
			fmt.Printf("%s\n", container.ID[:12])
		}
	}
}

func try(err error) {
	if err != nil {
		panic(err)
	}
}
