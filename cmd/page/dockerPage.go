package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/docker"
	"github.com/bbabos/swarmon/cmd/utils"
)

func dockerPage() {
	var selected string
	p := page{
		title: "DOCKER MENU",
		menuItems: []string{
			"1. Service options",
			"2. List containers",
			"3. List swarm nodes",
			"0. Back",
		},
	}
	renderPage(&p)

loop:
	for {
		fmt.Print("Select an option: ")
		selected = utils.ReadInput()

		switch selected {
		case "1":
			serviceOptions()
			break loop
		case "2":
			p.addSeparator()
			docker.ListContainers()
			fmt.Println(p.border)
		case "3":
			p.addSeparator()
			docker.ListSwarmNodes()
			fmt.Println(p.border)
		case "0":
			MenuPage()
			break loop
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}

func serviceOptions() {
	var selected string
	p := page{
		title:   "SERVICES",
		options: "| 0 - Back | 1 - another opts | 2 - another opts |",
		action:  docker.ListServices,
	}
	renderSubPage(&p)

loop:
	for {
		fmt.Print("Select an option: ")
		selected = utils.ReadInput()

		switch selected {
		case "1":
			// TODO
		case "2":
			// TODO
		case "0":
			dockerPage()
			break loop
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}
