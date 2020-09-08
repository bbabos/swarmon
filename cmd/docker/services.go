package docker

import (
	"context"
	"fmt"
	"log"

	"github.com/bbabos/swarmon/cmd/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// IService is ...
type IService interface {
	Inspect()
	Restart()
	GetName() string
}

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
		log.Fatal(err)
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

// Restart is ...
func (s Service) Restart() {
	fmt.Println("Service restart started...")
	utils.ExecShellCommand("docker service update --force "+s.ID, false)
	fmt.Println("Service restarted successfully.")
}

// Inspect is ...
func (s Service) Inspect() {
	utils.ExecShellCommand("docker service inspect "+s.ID, true)
}

// GetName is ...
func (s Service) GetName() string {
	return s.Name
}
