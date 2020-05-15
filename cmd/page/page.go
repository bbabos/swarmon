package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/utils"
	"github.com/manifoldco/promptui"
)

type page struct {
	Name   string
	action func()
}

func renderMenu(items []page, title string) {
	utils.Clear()
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\u2192 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: "{{ .Name }}"}

	prompt := promptui.Select{
		Label:     title,
		Items:     items,
		Templates: templates,
		Size:      5,
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
