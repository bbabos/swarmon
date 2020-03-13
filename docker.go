package main

import (
	"context"
	"fmt"
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
