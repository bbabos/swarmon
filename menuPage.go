package main

import "fmt"

func menuPage() {
	var selected string
	exit := false
	items := []string{
		"1. Monitoring stack options",
		"2. Maintain docker services",
		"0. Exit",
	}
	renderMenu(items, "MAIN MENU")

	for !exit {
		fmt.Println(border)
		fmt.Print("Select an option: ")
		selected = readInput()

		switch selected {
		case "1":
			exit = true
			stackPage()
		case "2":
			exit = true
			dockerPage()
		case "0":
			exit = true
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}

func dockerPage() {
	var selected string
	exit := false
	items := []string{
		"1. Container options",
		"2. Service options",
		"3. List swarm nodes",
		"0. Back",
	}
	renderMenu(items, "DOCKER MENU")

	for !exit {
		fmt.Println(border)
		fmt.Print("Select an option: ")
		selected = readInput()

		switch selected {
		case "1":
			exit = true
			containerOptions()
		case "2":
			listServices()
		case "3":
			listSwarmNodes()
		case "0":
			exit = true
			menuPage()
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}

func stackPage() {
	var selected string
	exit := false
	items := []string{
		"1. Docker stack deploy/update",
		"2. Remove monitoring stack",
		"0. Back",
	}
	renderMenu(items, "STACK MENU")

	for !exit {
		fmt.Println(border)
		fmt.Print("Select an option: ")
		selected = readInput()

		switch selected {
		case "1":
			stackInit()
		case "2":
			stackDelete()
		case "0":
			exit = true
			menuPage()
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}
