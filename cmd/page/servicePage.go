package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/docker"
)

func servicePage() {
	services := docker.GetServices()
	page := dynamicPage{
		details: `
--------- Service ----------
{{ "ID:" | faint }}	{{ .ID }}
{{ "Mode:" | faint }}	{{ .Mode }}
{{ "Replicas:" | faint }}	{{ .Replicas }}
{{ "CreatedAt:" | faint }}	{{ .Created }}
{{ "UpdatedAt:" | faint }}	{{ .Updated }}`,
		title: "SERVICES",
		items: services,
		size:  5,
	}
	selected := page.render()
	renderServiceSubPage(services[selected])
	defer dockerPage()
}

func renderServiceSubPage(s docker.IService) {
	sName := s.GetName()
	page := mainPage{
		title: sName,
		size:  5,
		items: []options{
			{Name: "Restart service", action: s.Restart},
			{Name: "Inspect service", action: s.Inspect},
			{Name: "Back", action: func() { return }},
		},
	}
	selected := page.render()
	page.items[selected].action()
	defer fmt.Println("----------------------------------------------")
}
