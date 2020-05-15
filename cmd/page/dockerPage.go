package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/docker"
	"github.com/bbabos/swarmon/cmd/utils"
)

func dockerPage() {
	var selected string
	var p iPage
	p = &menuPage{
		title: "DOCKER MENU",
		menuItems: []string{
			"1. Service options",
			"2. List containers",
			"3. List swarm nodes",
			"0. Back",
		},
	}
	p.renderMenuPage()

loop:
	for {
		fmt.Print("Select an option: ")
		selected = utils.ReadInput()

		switch selected {
		case "1":
			serviceOptions()
			break loop
		case "2":
			fmt.Println("---------TODO TITLE----------")
			docker.ListContainers()
			fmt.Println("---------TODO FOOTER----------")
		case "3":
			fmt.Println("---------TODO TITLE----------")
			docker.ListSwarmNodes()
			fmt.Println("---------TODO FOOTER----------")
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
