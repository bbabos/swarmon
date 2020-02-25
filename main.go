package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// var params []string

	fmt.Print("Enter your domain name (default: localhost): ")
	input, err := reader.ReadString('\n')
	try(err)
	fmt.Print(input)
}

func printContainersWithID() {
	cli, err := client.NewEnvClient()
	try(err)
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	try(err)
	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:12], container.Image)
	}
}

func try(err error) {
	if err != nil {
		panic(err)
	}
}
