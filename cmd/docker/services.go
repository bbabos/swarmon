package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
)

// Service is ...
type Service struct {
	ID       string
	Name     string
	Mode     string
	Replicas uint64
	Created  string
	Updated  string
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
			srv[i].Mode = "Replicated"
			srv[i].Replicas = *service.Spec.Mode.Replicated.Replicas
		} else {
			srv[i].Mode = "Global"

		}
		srv[i].ID = service.ID
		srv[i].Name = service.Spec.Name
		srv[i].Updated = service.UpdatedAt.Format("2006-01-02 15:04:05")
		srv[i].Created = service.CreatedAt.Format("2006-01-02 15:04:05")
	}
	return srv
}

// RestartService is ...
func RestartService(s Service) {
	cli, err := client.NewEnvClient()
	options := types.ServiceUpdateOptions{}
	spec := swarm.ServiceSpec{TaskTemplate: swarm.TaskSpec{ForceUpdate: 0}}

	_, err = cli.ServiceUpdate(context.Background(), s.ID, swarm.Version{}, spec, options)
	if err != nil {
		panic(err)
	}
}

// ScaleService is ...
func ScaleService(s Service) {
	fmt.Println("TODO")
}
