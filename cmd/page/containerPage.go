package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/docker"
	"github.com/manifoldco/promptui"
)

type containerOption struct {
	Name   string
	Action func(s docker.Container)
}

func containerPage() {
	containers := docker.GetContainers()
	renderContainerPage(containers)
}

func renderContainerPage(containers []docker.Container) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\u2192 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | white }}",
		Details: `
--------- Container ----------
{{ "ID:" | faint }}	{{ .ID }}
{{ "Created:" | faint }}	{{ .Created }}
{{ "State:" | faint }}	{{ .State }}`,
	}

	prompt := promptui.Select{
		Label:        "CONTAINERS",
		Items:        containers,
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
	renderContainerSubPage(containers[i])
}

func renderContainerSubPage(s docker.Container) {
	options := []containerOption{
		{Name: "Print container logs", Action: docker.GetContainerLogs},
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
	}
}
