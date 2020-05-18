package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/manifoldco/promptui"
)

// Service is ...
type Service struct {
	ID       string
	Name     string
	Mode     string
	Replicas uint64
}

type serviceOptions struct {
	Name   string
	Action func()
}

var options = []serviceOptions{
	{Name: "Restart service"},
	{Name: "Scale service"},
}

func (s *Service) restartService() {

}

// ServiceOptions is ...
func (s *Service) ServiceOptions() {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\u2192 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | white }}",
	}

	prompt := promptui.Select{
		Label:        s.Name,
		Items:        options,
		Templates:    templates,
		Size:         5,
		HideSelected: true,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	options[i].Action()
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
