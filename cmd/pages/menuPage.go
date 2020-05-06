package page

import "fmt"

func menuPage() {
	var selected string
	items := []string{
		"1. Monitoring stack options",
		"2. Maintain monitor services",
		"0. Exit",
	}
	renderMenu(items, "MAIN MENU")

loop:
	for {
		fmt.Print("Select an option: ")
		selected = readInput()

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
		selected = readInput()

		switch selected {
		case "1":
			serviceOptions()
			break loop
		case "2":
			listContainers()
		case "3":
			listSwarmNodes()
		case "0":
			menuPage()
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
	renderMenu(items, "STACK MENU")

loop:
	for {
		fmt.Print("Select an option: ")
		selected = readInput()

		switch selected {
		case "1":
			stackInit()
			break loop
		case "2":
			stackDelete()
		case "0":
			menuPage()
			break loop
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}
