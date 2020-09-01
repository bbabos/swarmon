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
	if i > 0 {
		renderNodeSubPage(nodes[i])
	}
	defer dockerPage()
}

func renderNodeSubPage(s docker.Node) {
	options := []nodeOption{
		{Name: "Promote node", Action: docker.Promote},
		{Name: "Demote node", Action: docker.Demote},
		{Name: "Back", Action: func(docker.Node) { return }},
	}
	i := renderPage(options, s.Name, "", 5)
	options[i].Action(s)
}
