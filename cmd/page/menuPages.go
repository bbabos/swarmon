package page

// MainPage is ...
func MainPage() {
	options := []options{
		{Name: "Monitoring stack", action: stackPage},
		{Name: "Docker options", action: dockerPage},
		{Name: "Exit", action: func() { return }},
	}
	page := page{
		details: "",
		title:   "MAIN MENU",
		items:   options,
		size:    5,
	}
	selected := page.render()
	options[selected].action()
}

func dockerPage() {
	options := []options{
		{Name: "Services", action: servicePage},
		{Name: "Containers", action: containerPage},
		{Name: "Nodes", action: nodePage},
		{Name: "Back", action: MainPage},
	}
	page := page{
		details: "",
		title:   "DOCKER MENU",
		items:   options,
		size:    5,
	}
	selected := page.render()
	options[selected].action()
}

func stackPage() {
	options := []options{
		{Name: "Docker stack deploy/update", action: stackInitOrUpdate},
		{Name: "Remove monitoring stack", action: stackDelete},
		{Name: "Back", action: MainPage},
	}
	page := page{
		details: "",
		title:   "STACK MENU",
		items:   options,
		size:    5,
	}
	selected := page.render()
	options[selected].action()
}
