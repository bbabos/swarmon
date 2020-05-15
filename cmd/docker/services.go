package docker

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Service is ...
type Service struct {
	ID       string
	Name     string
	Mode     string
	Replicas uint64
	Action   func()
}

// GetServices is ...
func GetServices() []Service {
	cli, err := client.NewEnvClient()
	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		panic(err)
	}

	srv := make([]Service, len(services))

	for i, service := range services {
		if service.Spec.Mode.Replicated != nil {
			srv[i].ID = service.ID
			srv[i].Name = service.Spec.Name
			srv[i].Mode = "Replicated"
			srv[i].Replicas = *service.Spec.Mode.Replicated.Replicas
		} else {
			srv[i].ID = service.ID
			srv[i].Name = service.Spec.Name
			srv[i].Mode = "Global"
		}
	}
	return srv
}

// GetLogs is ...
func (s Service) GetLogs() {
	cli, err := client.NewEnvClient()
	logs, err := cli.ServiceLogs(context.Background(), s.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	p := make([]byte, 256)
	for {
		n, err := logs.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Print(string(p[:n]))
				break
			}
			fmt.Println(err)
		}
		fmt.Print(string(p))
	}
}
