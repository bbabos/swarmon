package page

// MenuPage is ...
func MenuPage() {
	p := []page{
		{Name: "Monitoring stack options", action: stackPage},
		{Name: "Maintain monitor services", action: dockerPage},
		{Name: "Exit"},
	}
	renderMenu(p, "MAIN MENU")
}
