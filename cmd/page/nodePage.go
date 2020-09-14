package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/docker"
)

func nodePage() {
	nodes := docker.GetNodes()
	mainPage := page{
		details: `
--------- node ----------
{{ "ID:" | faint }}	{{ .ID }}
{{ "Status:" | faint }}	{{ .Status }}
{{ "Role:" | faint }}	{{ .Role }}
{{ "Availability:" | faint }}	{{ .Availability }}
{{ "EngineVersion:" | faint }}	{{ .EngineVersion }}`,
		title: "NODES",
		items: nodes,
		size:  5,
	}
	selected := mainPage.render()
	renderNodeSubPage(nodes[selected])
	defer dockerPage()
}

func renderNodeSubPage(n docker.INode) {
	options := []options{
		{Name: "Promote node", action: n.Promote},
		{Name: "Demote node", action: n.Demote},
		{Name: "Back", action: func() { return }},
	}
	nName := n.GetName()
	subPage := page{
		details: "",
		title:   nName,
		items:   options,
		size:    5,
	}
	selected := subPage.render()
	options[selected].action()
	defer fmt.Println("----------------------------------------------")
}
