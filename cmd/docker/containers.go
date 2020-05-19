package docker

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

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
		panic(err)
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

// GetContainerLogs is ...
func GetContainerLogs(c Container) {
	cli, err := client.NewEnvClient()
	out, err := cli.ContainerLogs(context.Background(), c.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)
}
