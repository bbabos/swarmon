package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/docker"
)

func servicePage() {
	services := docker.GetServices()
	mainPage := page{
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
	selected := mainPage.render()
	renderServiceSubPage(services[selected])
	defer dockerPage()
}

func renderServiceSubPage(s docker.IService) {
	options := []options{
		{Name: "Restart service", action: s.Restart},
		{Name: "Inspect service", action: s.Inspect},
		{Name: "Back", action: func() { return }},
	}
	sName := s.GetName()
	subPage := page{
		details: "",
		title:   sName,
		items:   options,
		size:    5,
	}
	selected := subPage.render()
	options[selected].action()
	defer fmt.Println("----------------------------------------------")
}
