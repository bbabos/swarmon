package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var stackFile = "tmp/parsed.yml"

func listContainers() {
	cli, err := client.NewEnvClient()
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	try(err)

	fmt.Println("--------------------------")
	fmt.Println("CONTAINERS:")
	fmt.Println("--------------------------")
	for _, container := range containers {
		fmt.Printf("%s | %s\n", container.ID[:12], container.Names)
	}
	fmt.Println("--------------------------")

}

func listServices() {
	cli, err := client.NewEnvClient()
	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	try(err)

	fmt.Println("--------------------------")
	fmt.Println("SWARM SERVICES:           |")
	fmt.Println("--------------------------")
	for _, service := range services {
		fmt.Printf("%s | %s\n", service.ID, service.Spec.Name)
	}
	fmt.Println("--------------------------")
}

func listSwarmNodes() {
	cli, err := client.NewEnvClient()
	nodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	try(err)

	fmt.Println("--------------------------")
	fmt.Println("SWARM NODES:              |")
	fmt.Println("--------------------------")
	for _, node := range nodes {
		fmt.Printf("%s | %s | %s | %s\n", node.ID, node.Description.Hostname, node.Spec.Role, node.Status.State)
	}
	fmt.Println("--------------------------")
}

func deployStack() {
	command := "docker stack deploy -c " + stackFile + " " + p.Docker.StackName
	execCommand(command)
}

func removeStack() {
	command := "docker stack rm " + p.Docker.StackName
	execCommand(command)
}
