package page

import (
	"fmt"

	"github.com/bbabos/swarmon-go/cmd/docker"
	"github.com/bbabos/swarmon-go/cmd/utils"
)

// MenuPage is ...
func MenuPage() {
	var selected string
	items := []string{
		"1. Monitoring stack options",
		"2. Maintain monitor services",
		"0. Exit",
	}
	RenderMenu(items, "MAIN MENU")

loop:
	for {
		fmt.Print("Select an option: ")
		selected = utils.ReadInput()

		switch selected {
		case "1":
			stackPage()
			break loop
		case "2":
			DockerPage()
			break loop
		case "0":
			break loop
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}

// DockerPage is ...
func DockerPage() {
	var selected string
	items := []string{
		"1. Service options",
		"2. List containers",
		"3. List swarm nodes",
		"0. Back",
	}
	RenderMenu(items, "DOCKER MENU")

loop:
	for {
		fmt.Print("Select an option: ")
		selected = utils.ReadInput()

		switch selected {
		case "1":
			ServiceOptions()
			break loop
		case "2":
			docker.ListContainers()
		case "3":
			docker.ListSwarmNodes()
		case "0":
			MenuPage()
			break loop
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}

func stackPage() {
	var selected string
	items := []string{
		"1. Docker stack deploy/update",
		"2. Remove monitoring stack",
		"0. Back",
	}
	RenderMenu(items, "STACK MENU")

loop:
	for {
		fmt.Print("Select an option: ")
		selected = utils.ReadInput()

		switch selected {
		case "1":
			StackInit()
			break loop
		case "2":
			StackDelete()
		case "0":
			MenuPage()
			break loop
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}
