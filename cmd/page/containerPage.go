package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/docker"
)

func containerPage() {
	containers := docker.GetContainers()
	page := dynamicPage{
		details: `
--------- Container ----------
{{ "ID:" | faint }}	{{ .ID }}
{{ "Created:" | faint }}	{{ .Created }}
{{ "State:" | faint }}	{{ .State }}`,
		title: "CONTAINERS",
		items: containers,
		size:  10,
	}
	selected := page.render()
	renderContainerSubPage(containers[selected])
	defer dockerPage()
}

func renderContainerSubPage(c docker.IContainer) {
	cName := c.GetName()
	page := mainPage{
		title: cName,
		size:  5,
		items: []options{
			{Name: "Print logs", action: c.GetLogs},
			{Name: "Inspect", action: c.Inspect},
			{Name: "Back", action: func() { return }},
		},
	}
	selected := page.render()
	page.items[selected].action()
	defer fmt.Println("----------------------------------------------")
}
