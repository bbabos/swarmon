package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/docker"
	"github.com/bbabos/swarmon/cmd/utils"
)

func dockerPage() {
	p := []page{
		{Name: "Service options", action: serviceOptions},
		{Name: "Container options", action: containerOptions},
		{Name: "Node options", action: nodeOptions},
		{Name: "Exit"},
	}
	renderMenu(p, "DOCKER MENU")
}

func containerOptions() {
	fmt.Println("TODO")
}

func nodeOptions() {
	fmt.Println("TODO")
}

func serviceOptions() {
	var selected string
	var page iSubPage
	page = &subPage{
		title:   "SERVICES",
		options: "| 0 - Back | 1 - Scale | 2 - Restart |",
		action:  docker.ListServices,
	}
	page.renderSubPage()

loop:
	for {
		fmt.Print("Select an option: ")
		selected = utils.ReadInput()

		switch selected {
		case "1":
			// TODO
			fmt.Println("TODO")
		case "2":
			// TODO
			fmt.Println("TODO")
		case "0":
			dockerPage()
			break loop
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}
