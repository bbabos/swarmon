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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter image name: ")
	var input string
	if scanner.Scan() {
		input = scanner.Text()
	}

	printContainerID(input)
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
