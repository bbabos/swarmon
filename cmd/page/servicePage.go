package page

import (
	"github.com/bbabos/swarmon/cmd/docker"
)

type serviceOption struct {
	Name   string
	Action func(s docker.Service)
}

func servicePage() {
	services := docker.GetServices()
	details := `
--------- Service ----------
{{ "ID:" | faint }}	{{ .ID }}
{{ "Mode:" | faint }}	{{ .Mode }}
{{ "Replicas:" | faint }}	{{ .Replicas }}
{{ "CreatedAt:" | faint }}	{{ .Created }}
{{ "UpdatedAt:" | faint }}	{{ .Updated }}`
	i := renderPage(services, "SERVICES", details, 5)
	renderServiceSubPage(services[i])
}

func renderServiceSubPage(s docker.Service) {
	options := []serviceOption{
		{Name: "Restart service", Action: docker.RestartService},
		{Name: "Scale service", Action: docker.ScaleService},
		{Name: "Back", Action: func(docker.Service) { return }},
	}

	i := renderPage(options, s.Name, "", 5)
	options[i].Action(s)
}
