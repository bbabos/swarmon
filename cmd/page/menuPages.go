package page

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// MainPage is ...
func MainPage() {
	p := []page{
		{Name: "Monitoring stack options", action: stackPage},
		{Name: "Swarm options", action: dockerPage},
		{Name: "Exit"},
	}
	renderMenu(p, "MAIN MENU")
}

func dockerPage() {
	p := []page{
		{Name: "Services", action: servicePage},
		{Name: "Containers", action: containerPage},
		{Name: "Nodes", action: nodePage},
		{Name: "Back", action: MainPage},
	}
	renderMenu(p, "DOCKER MENU")
}

type page struct {
	Name   string
	action func()
}

func renderMenu(items []page, title string) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\u2192 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | white }}"}

	prompt := promptui.Select{
		Label:        title,
		Items:        items,
		Templates:    templates,
		Size:         5,
		HideSelected: true,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if items[i].Name == "Exit" {
		return
	}
	items[i].action()
}
