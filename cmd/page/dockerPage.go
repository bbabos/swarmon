package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/docker"
	"github.com/manifoldco/promptui"
)

func dockerPage() {
	p := []page{
		{Name: "Service options", action: serviceOptions},
		{Name: "Container options", action: containerOptions},
		{Name: "Node options", action: nodeOptions},
		{Name: "Back", action: MenuPage},
	}
	renderMenu(p, "DOCKER MENU")
}

func serviceOptions() {
	services := docker.GetServices()
	renderService(services)
}

func renderService(services []docker.Service) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\u2192 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: "{{ .Name | cyan }}",
		Details: `
--------- Service ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Mode:" | faint }}	{{ .Mode }}
{{ "Replicas:" | faint }}	{{ .Replicas }}`,
	}

	prompt := promptui.Select{
		Label:     "SERVICES",
		Items:     services,
		Templates: templates,
		Size:      10,
	}

	_, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// services[i].Action()
}

func containerOptions() {
	fmt.Println("TODO")
}

func nodeOptions() {
	fmt.Println("TODO")
}
