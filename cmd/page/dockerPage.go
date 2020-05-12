package page

import (
	"fmt"

	"github.com/bbabos/swarmon-go/cmd/docker"
	"github.com/bbabos/swarmon-go/cmd/utils"
)

func dockerPage() {
	var selected string
	items := []string{
		"1. Service options",
		"2. List containers",
		"3. List swarm nodes",
		"0. Back",
	}
	renderMenu(items, "DOCKER MENU")

loop:
	for {
		fmt.Print("Select an option: ")
		selected = utils.ReadInput()

		switch selected {
		case "1":
			serviceOptions()
			break loop
		case "2":
			fmt.Println(border)
			fmt.Println("CONTAINERS")
			fmt.Println(border)
			docker.ListContainers()
			fmt.Println(border)
		case "3":
			fmt.Println(border)
			fmt.Println("SWARM NODES")
			fmt.Println(border)
			docker.ListSwarmNodes()
			fmt.Println(border)
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
	pageItem := Page{
		Title:   "SERVICES",
		Options: "| 0 - Back | 1 - another opts | 2 - another opts |",
		Action:  docker.ListServices,
	}
	renderPage(pageItem)

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
