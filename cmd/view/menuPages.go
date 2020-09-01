package page

// MainPage is ...
func MainPage() {
	p := []page{
		{Name: "Monitoring stack", action: stackPage},
		{Name: "Docker options", action: dockerPage},
		{Name: "Exit", action: func() { return }},
	}
	i := renderPage(p, "MAIN MENU", "", 5)
	p[i].action()
}

func dockerPage() {
	p := []page{
		{Name: "Services", action: servicePage},
		{Name: "Containers", action: containerPage},
		{Name: "Nodes", action: nodePage},
		{Name: "Back", action: MainPage},
	}
	i := renderPage(p, "DOCKER MENU", "", 5)
	p[i].action()
}

func stackPage() {
	p := []page{
		{Name: "Docker stack deploy/update", action: stackInitOrUpdate},
		{Name: "Remove monitoring stack", action: stackDelete},
		{Name: "Back", action: MainPage},
	}
	i := renderPage(p, "STACK MENU", "", 5)
	p[i].action()
}