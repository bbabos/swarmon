package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/docker"
)

func containerPage() {
	containers := docker.GetContainers()
	mainPage := page{
		details: `
--------- Container ----------
{{ "ID:" | faint }}	{{ .ID }}
{{ "Created:" | faint }}	{{ .Created }}
{{ "State:" | faint }}	{{ .State }}`,
		title: "CONTAINERS",
		items: containers,
		size:  10,
	}
	selected := mainPage.render()
	renderContainerSubPage(containers[selected])
	defer dockerPage()
}

func renderContainerSubPage(c docker.IContainer) {
	options := []options{
		{Name: "Print logs", action: c.GetLogs},
		{Name: "Inspect", action: c.Inspect},
		{Name: "Back", action: func() { return }},
	}
	cName := c.GetName()
	subPage := page{
		details: "",
		title:   cName,
		items:   options,
		size:    5,
	}
	selected := subPage.render()
	options[selected].action()
	defer fmt.Println("----------------------------------------------")
}
