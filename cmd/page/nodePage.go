package page

import (
	"github.com/bbabos/swarmon/cmd/docker"
)

type nodeOption struct {
	Name   string
	Action func(n docker.Node)
}

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
}

func renderNodeSubPage(s docker.Node) {
	options := []nodeOption{
		{Name: "Promote node", Action: docker.PromoteNode},
		{Name: "Demote node", Action: docker.DemoteNode},
		{Name: "Back", Action: func(docker.Node) { return }},
	}

	i := renderPage(options, s.Name, "", 5)
	options[i].Action(s)
}
