package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func listContainerIDs() {
	cli, err := client.NewEnvClient()
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	try(err)

	fmt.Println("\n--------------------------")
	fmt.Println("CONTAINERS:")
	for _, container := range containers {
		fmt.Printf("%s\n", container.ID[:12])
	}
}

func listServices() {
	cli, err := client.NewEnvClient()
	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	try(err)

	fmt.Println("\n--------------------------")
	fmt.Println("SWARM SERVICES:")
	for _, service := range services {
		fmt.Printf("%s | %s\n", service.ID, service.Spec.Name)
	}
}

func listSwarmNodes() {
	cli, err := client.NewEnvClient()
	nodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	try(err)

	fmt.Println("\n--------------------------")
	fmt.Println("SWARM NODES:")
	for _, node := range nodes {
		fmt.Printf("%s | %s | %s | %s\n", node.ID, node.Description.Hostname, node.Spec.Role, node.Status.State)
	}
}

func stackDeploy(stackFile string, stackName string) {
	cmd := exec.Command("docker", "stack", "deploy", "-c", stackFile, stackName)
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Run()

	if stderr.Len() != 0 {
		fmt.Println()
		fmt.Printf("Error: %v", stderr.String())
	} else {
		fmt.Println()
		fmt.Printf("Result:\n%v", out.String())
	}
}
