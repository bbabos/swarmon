package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func listContainerIDbasedOnImage(cli *client.Client, imageName string) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	try(err)

	fmt.Println("\n--------------------------")
	fmt.Println("CONTAINERS\n--------------------------")
	for _, container := range containers {
		if strings.Contains(container.Image, imageName) {
			fmt.Printf("%s\n", container.ID[:12])
		}
	}
}

func listServices(cli *client.Client) {
	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	try(err)

	fmt.Println("\n--------------------------")
	fmt.Println("SWARM SERVICES\n--------------------------")
	for _, service := range services {
		fmt.Printf("%s | %s", service.ID, service.Spec.Name)
	}
}

func listSwarmNodes(cli *client.Client) {
	nodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	try(err)

	fmt.Println("\n--------------------------")
	fmt.Println("SWARM NODES\n--------------------------")
	for _, node := range nodes {
		fmt.Printf("%s | %s | %s | %s\n", node.ID, node.Description.Hostname, node.Spec.Role, node.Status.State)
	}
}
