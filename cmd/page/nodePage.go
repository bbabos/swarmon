package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/docker"
)

func nodePage() {
	nodes := docker.GetNodes()
	page := dynamicPage{
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
	selected := page.render()
	renderNodeSubPage(nodes[selected])
	defer dockerPage()
}

func renderNodeSubPage(n docker.INode) {
	nName := n.GetName()
	page := mainPage{
		title: nName,
		size:  5,
		items: []options{
			{Name: "Promote node", action: n.Promote},
			{Name: "Demote node", action: n.Demote},
			{Name: "Back", action: func() { return }},
		},
	}
	selected := page.render()
	page.items[selected].action()
	defer fmt.Println("----------------------------------------------")
}
