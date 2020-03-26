package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func listContainers() {
	cli, err := client.NewEnvClient()
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	try(err)

	fmt.Println(border)
	fmt.Println("CONTAINERS")
	fmt.Println(border)
	for _, container := range containers {
		fmt.Printf("%s | %s\n", container.ID[:12], container.Names)
	}
}

func listServices() {
	cli, err := client.NewEnvClient()
	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	try(err)

	fmt.Println(border)
	fmt.Println("SERVICES")
	fmt.Println(border)
	for _, service := range services {
		fmt.Printf("%s | %s\n", service.ID, service.Spec.Name)
	}
}

func listSwarmNodes() {
	cli, err := client.NewEnvClient()
	nodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	try(err)

	fmt.Println(border)
	fmt.Println("SWARM NODES")
	fmt.Println(border)
	for _, node := range nodes {
		fmt.Printf("%s | %s | %s | %s\n", node.ID, node.Description.Hostname, node.Spec.Role, node.Status.State)
	}
}
