package page

// MainPage is ...
func MainPage() {
	page := mainPage{
		title: "MAIN MENU",
		size:  5,
		items: []options{
			{Name: "Monitoring stack", action: stackPage},
			{Name: "Docker options", action: dockerPage},
			{Name: "Exit", action: func() { return }},
		},
	}
	selected := page.render()
	page.items[selected].action()
}

func dockerPage() {
	page := mainPage{
		title: "DOCKER MENU",
		size:  5,
		items: []options{
			{Name: "Services", action: servicePage},
			{Name: "Containers", action: containerPage},
			{Name: "Nodes", action: nodePage},
			{Name: "Back", action: MainPage},
		},
	}
	selected := page.render()
	page.items[selected].action()
}

func stackPage() {
	page := mainPage{
		title: "STACK MENU",
		size:  5,
		items: []options{
			{Name: "Docker stack deploy/update", action: stackInitOrUpdate},
			{Name: "Remove monitoring stack", action: stackDelete},
			{Name: "Back", action: MainPage},
		},
	}
	selected := page.render()
	page.items[selected].action()
}
