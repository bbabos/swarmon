package main

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

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

// TODO with SDK
func deployStack() {
	exec.Command("docker", "stack deploy -c ./templates/compose.yml swarmon")
}
