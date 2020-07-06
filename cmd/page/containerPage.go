package page

import (
	"github.com/bbabos/swarmon/cmd/docker"
)

type containerOption struct {
	Name   string
	Action func(s docker.Container)
}

func containerPage() {
	containers := docker.GetContainers()
	details := `
--------- Container ----------
{{ "ID:" | faint }}	{{ .ID }}
{{ "Created:" | faint }}	{{ .Created }}
{{ "State:" | faint }}	{{ .State }}`
	i := renderPage(containers, "CONTAINERS", details, 10)
	renderContainerSubPage(containers[i])
}

func renderContainerSubPage(s docker.Container) {
	options := []containerOption{
		{Name: "Print container logs", Action: docker.GetContainerLogs},
		{Name: "Back", Action: func(docker.Container) { return }},
	}
	defer dockerPage()

	i := renderPage(options, s.Name, "", 5)
	options[i].Action(s)
}
