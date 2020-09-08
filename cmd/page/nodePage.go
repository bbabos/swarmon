package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/docker"
)

func nodePage() {
	nodes := docker.GetNodes()
	details := `
--------- node ----------
{{ "ID:" | faint }}	{{ .ID }}
{{ "Status:" | faint }}	{{ .Status }}
{{ "Role:" | faint }}	{{ .Role }}
{{ "Availability:" | faint }}	{{ .Availability }}
{{ "EngineVersion:" | faint }}	{{ .EngineVersion }}`
	i := renderPage(nodes, "NODES", details, 5)
	renderNodeSubPage(nodes[i])
	defer dockerPage()
}

func renderNodeSubPage(n docker.INode) {
	options := []page{
		{Name: "Promote node", action: n.Promote},
		{Name: "Demote node", action: n.Demote},
		{Name: "Back", action: func() { return }},
	}
	nName := n.GetName()
	i := renderPage(options, nName, "", 5)
	options[i].action()
	defer fmt.Println("----------------------------------------------")
}
