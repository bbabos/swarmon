package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/docker"
)

func containerPage() {
	containers := docker.GetContainers()
	details := `
--------- Container ----------
{{ "ID:" | faint }}	{{ .ID }}
{{ "Created:" | faint }}	{{ .Created }}
{{ "State:" | faint }}	{{ .State }}`
	i := renderPage(containers, "CONTAINERS", details, 10)
	renderContainerSubPage(containers[i])
	defer dockerPage()
}

func renderContainerSubPage(c docker.IContainer) {
	options := []page{
		{Name: "Print logs", action: c.GetLogs},
		{Name: "Inspect", action: c.Inspect},
		{Name: "Back", action: func() { return }},
	}
	cName := c.GetName()
	i := renderPage(options, cName, "", 5)

	options[i].action()
	defer fmt.Println("----------------------------------------------")
}
