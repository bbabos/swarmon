package docker

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Node is ...
type Node struct {
	ID            string
	Name          string
	Status        string
	Availability  string
	EngineVersion string
	Role          string
}

// GetNodes is ...
func GetNodes() []Node {
	cli, err := client.NewEnvClient()
	nodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	nds := make([]Node, len(nodes))

	for i, node := range nodes {
		nds[i].ID = node.ID
		nds[i].Name = node.Description.Hostname
		nds[i].Status = string(node.Status.State)
		nds[i].Availability = string(node.Spec.Availability)
		nds[i].EngineVersion = node.Description.Engine.EngineVersion
		nds[i].Role = string(node.Spec.Role)
	}
	return nds
}

// PromoteNode is ...
func PromoteNode(n Node) {
	// TODO
}

// DemoteNode is ...
func DemoteNode(n Node) {
	// TODO
}
