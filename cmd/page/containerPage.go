package page

import (
	"github.com/bbabos/swarmon/cmd/docker"
)

func containerPage() {
	containers := docker.GetContainers()
	page := dynamicPage{
		details: `
--------- Info ----------
{{ "ID:" | faint }}	{{ .ID }}
{{ "Created:" | faint }}	{{ .Created }}
{{ "State:" | faint }}	{{ .State }}`,
		title: "CONTAINERS",
		items: containers,
		size:  10,
	}
	selected := page.renderBase()
	containerSubPage(containers[selected])
	defer dockerPage()
}

func containerSubPage(c docker.IContainer) {
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
	page.render()
}
