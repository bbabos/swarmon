package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/docker"
	"github.com/manifoldco/promptui"
)

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
{{ "Replicas:" | faint }}	{{ .Replicas }}`,
	}

	prompt := promptui.Select{
		Label:        "SERVICES",
		Items:        services,
		Templates:    templates,
		Size:         10,
		HideSelected: true,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	services[i].ServiceOptions()
}
