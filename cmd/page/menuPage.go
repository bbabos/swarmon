package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/utils"
)

// MenuPage is ...
func MenuPage() {
	var selected string
	p := page{
		title: "MAIN MENU",
		menuItems: []string{
			"1. Monitoring stack options",
			"2. Maintain monitor services",
			"0. Exit",
		},
	}
	renderPage(&p)

loop:
	for {
		fmt.Print("Select an option: ")
		selected = utils.ReadInput()

		switch selected {
		case "1":
			stackPage()
			break loop
		case "2":
			dockerPage()
			break loop
		case "0":
			break loop
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}
