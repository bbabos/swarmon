package page

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type page struct {
	Name   string
	action func()
}

func renderPage(items interface{}, title string, details string, size int) int {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\u2192 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | white }}",
		Details:  details,
	}

	prompt := promptui.Select{
		Label:        title,
		Items:        items,
		Templates:    templates,
		Size:         size,
		HideSelected: true,
		HideHelp:     true,
	}
	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
	return i
}
