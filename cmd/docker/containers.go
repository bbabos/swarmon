package docker

import (
	"context"
	"log"
	"time"

	"github.com/bbabos/swarmon/cmd/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// IContainer is ...
type IContainer interface {
	GetLogs()
	Inspect()
	GetName() string
}

// Container is ...
type Container struct {
	ID      string
	Name    string
	State   string
	Created string
}

// GetContainers is ...
func GetContainers() []Container {
	cli, err := client.NewEnvClient()
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	conts := make([]Container, len(containers))

	for i, container := range containers {
		conts[i].ID = container.ID[:12]
		conts[i].Name = container.Names[0]
		conts[i].State = container.State
		conts[i].Created = time.Unix(container.Created, 0).Format("2006-01-02 15:04:05")
	}
	return conts
}

// GetLogs is ...
func (c Container) GetLogs() {
	utils.ExecShellCommand("docker logs "+c.ID, true)
}

// Inspect is ...
func (c Container) Inspect() {
	utils.ExecShellCommand("docker inspect "+c.ID, true)
}

// GetName is ...
func (c Container) GetName() string {
	return c.Name
}
