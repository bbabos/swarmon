package page

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type page struct {
	title   string
	details string
	size    int
	items   interface{}
}

type options struct {
	Name   string
	action func()
}

func (p page) render() int {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\u2192 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | white }}",
		Details:  p.details,
	}

	prompt := promptui.Select{
		Label:        p.title,
		Items:        p.items,
		Templates:    templates,
		Size:         p.size,
		HideSelected: true,
		HideHelp:     true,
	}
	selected, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt exited %v\n", err)
	}
	return selected
}
