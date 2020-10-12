package page

import (
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
	selected := page.renderBase()
	nodeSubPage(nodes[selected])
	defer dockerPage()
}

func nodeSubPage(n docker.INode) {
	nName := n.GetName()
	page := mainPage{
		title: nName,
		size:  5,
		items: []options{
			{Name: "Promote node", action: n.Promote},
			{Name: "Demote node", action: n.Demote},
			{Name: "Inspect node", action: n.Inspect},
			{Name: "Back", action: func() { return }},
		},
	}
	page.render()
}
