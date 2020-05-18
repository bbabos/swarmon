package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/docker"
	"github.com/manifoldco/promptui"
)

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
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	containers[i].GetLogs()
}
