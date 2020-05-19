package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/docker"
	"github.com/manifoldco/promptui"
)

type nodeOption struct {
	Name   string
	Action func(n docker.Node)
}

func nodePage() {
	nodes := docker.GetNodes()
	renderNodePage(nodes)
}

func renderNodePage(nodes []docker.Node) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\u2192 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | white }}",
		Details: `
--------- node ----------
{{ "ID:" | faint }}	{{ .ID }}
{{ "Status:" | faint }}	{{ .Status }}
{{ "Role:" | faint }}	{{ .Role }}
{{ "Availability:" | faint }}	{{ .Availability }}
{{ "EngineVersion:" | faint }}	{{ .EngineVersion }}`,
	}

	prompt := promptui.Select{
		Label:        "NODES",
		Items:        nodes,
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

	renderNodeSubPage(nodes[i])
}

func renderNodeSubPage(s docker.Node) {
	options := []nodeOption{
		{Name: "Promote node", Action: docker.PromoteNode},
		{Name: "Demote node", Action: docker.DemoteNode},
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
