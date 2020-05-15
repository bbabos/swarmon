package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Service is ...
type Service struct {
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
			srv[i].Name = service.Spec.Name
			srv[i].Mode = "Replicated"
			srv[i].Replicas = *service.Spec.Mode.Replicated.Replicas
		} else {
			srv[i].Name = service.Spec.Name
			srv[i].Mode = "Global"
		}
	}
	return srv
}
