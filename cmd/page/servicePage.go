package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/docker"
)

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
	defer dockerPage()
}

func renderServiceSubPage(s docker.IService) {
	options := []page{
		{Name: "Restart service", action: s.Restart},
		{Name: "Inspect service", action: s.Inspect},
		{Name: "Back", action: func() { return }},
	}
	sName := s.GetName()
	i := renderPage(options, sName, "", 5)
	options[i].action()
	defer fmt.Println("----------------------------------------------")
}
