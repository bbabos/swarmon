package menupage

import (
	"fmt"

	"github.com/bbabos/swarmon-go/cmd/docker"
	"github.com/bbabos/swarmon-go/cmd/page"
	"github.com/bbabos/swarmon-go/cmd/page/dockerpage"
	"github.com/bbabos/swarmon-go/cmd/page/stackpage"
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
	page.RenderMenu(items, "MAIN MENU")

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
	page.RenderMenu(items, "DOCKER MENU")

loop:
	for {
		fmt.Print("Select an option: ")
		selected = utils.ReadInput()

		switch selected {
		case "1":
			dockerpage.ServiceOptions()
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
	page.RenderMenu(items, "STACK MENU")

loop:
	for {
		fmt.Print("Select an option: ")
		selected = utils.ReadInput()

		switch selected {
		case "1":
			stackpage.StackInit()
			break loop
		case "2":
			stackpage.StackDelete()
		case "0":
			MenuPage()
			break loop
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}
