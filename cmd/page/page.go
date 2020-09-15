package page

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type mainPage struct {
	title string
	size  int
	items []options
}

type dynamicPage struct {
	title   string
	details string
	size    int
	items   interface{}
}

type options struct {
	Name   string
	action func()
}

func (p mainPage) render() {
	selected := p.renderBase()
	p.runAction(selected)
}

func (p mainPage) runAction(selected int) func() {
	return p.items[selected].action
}

func (p mainPage) renderBase() int {
	prompt := promptui.Select{
		Label: p.title,
		Items: p.items,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "\u2192 {{ .Name | cyan }}",
			Inactive: "  {{ .Name | white }}",
		},
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

func (p dynamicPage) renderBase() int {
	prompt := promptui.Select{
		Label: p.title,
		Items: p.items,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "\u2192 {{ .Name | cyan }}",
			Inactive: "  {{ .Name | white }}",
			Details:  p.details,
		},
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
