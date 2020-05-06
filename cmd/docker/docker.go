package docker

import (
	"context"
	"fmt"

	"github.com/bbabos/swarmon-go/config"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// ListContainers is ...
func ListContainers() {
	cli, err := client.NewEnvClient()
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println(config.Border)
	fmt.Println("CONTAINERS")
	fmt.Println(config.Border)
	for _, container := range containers {
		fmt.Printf("%s | %s | %s\n", container.ID[:12], container.Status, container.Names)
	}
	fmt.Println(config.Border)
}

// ListServices is ...
func ListServices() {
	cli, err := client.NewEnvClient()
	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		panic(err)
	}

	for _, service := range services {
		fmt.Printf("%s | %s\n", service.ID, service.Spec.Name)
	}
}

// ListSwarmNodes is ...
func ListSwarmNodes() {
	cli, err := client.NewEnvClient()
	nodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println(config.Border)
	fmt.Println("SWARM NODES")
	fmt.Println(config.Border)
	for _, node := range nodes {
		fmt.Printf("%s | %s | %s | %s\n", node.ID, node.Description.Hostname, node.Spec.Role, node.Status.State)
	}
	fmt.Println(config.Border)
}
