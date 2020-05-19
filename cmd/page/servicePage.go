package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/docker"
	"github.com/manifoldco/promptui"
)

type serviceOptions struct {
	Name   string
	Action func(s docker.Service)
}

func servicePage() {
	services := docker.GetServices()
	renderServicePage(services)
}

func renderServicePage(services []docker.Service) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\u2192 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | white }}",
		Details: `
--------- Service ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Mode:" | faint }}	{{ .Mode }}
{{ "Replicas:" | faint }}	{{ .Replicas }}
{{ "CreatedAt:" | faint }}	{{ .Created }}
{{ "UpdatedAt:" | faint }}	{{ .Updated }}`,
	}

	prompt := promptui.Select{
		Label:        "SERVICES",
		Items:        services,
		Templates:    templates,
		Size:         10,
		HideSelected: true,
		HideHelp:     true,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	renderServicesSubPage(services[i])
}

func renderServicesSubPage(s docker.Service) {
	options := []serviceOptions{
		{Name: "Restart service", Action: docker.RestartService},
		{Name: "Scale service", Action: docker.ScaleService},
		{Name: "Back"},
	}

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
		HideHelp:     true,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if options[i].Name == "Back" {
		dockerPage()
	} else {
		options[i].Action(s)
		dockerPage()
	}
}
