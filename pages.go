package main

import "fmt"

func renderMenu(items []string, border string, title string) {
	clear()
	fmt.Println(border)
	fmt.Println(title)
	fmt.Println(border)
	for _, item := range items {
		fmt.Println(item)
	}
}

func createBorder(items []string) string {
	border := ""
	length := 0
	for _, item := range items {
		if len(item) > length {
			length = len(item)
		}
	}
	for i := 0; i < length; i++ {
		border += "-"
	}
	return border
}

func menuPage() {
	var selected string
	items := []string{
		"1. Monitoring stack options",
		"2. Maintain monitoring services",
		"0. Exit",
	}
	border := createBorder(items)
	renderMenu(items, border, "MAIN MENU")

	for selected != "0" {
		fmt.Println(border)
		fmt.Print("Choose an option: ")
		selected = readInput()

		switch selected {
		case "1":
			initPage()
			return
		case "2":
			dockerPage()
			return
		case "0":
			return
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}

func dockerPage() {
	var selected string
	items := []string{
		"1. List all running containers",
		"2. List services",
		"3. List swarm nodes",
		"0. Back",
	}
	border := createBorder(items)
	renderMenu(items, border, "DOCKER MENU")

	for selected != "0" {
		fmt.Println(border)
		fmt.Print("Choose an option: ")
		selected = readInput()

		switch selected {
		case "1":
			fmt.Println("--------------------------")
			fmt.Println("CONTAINERS:               |")
			fmt.Println("--------------------------")
			listContainers()
		case "2":
			fmt.Println("--------------------------")
			fmt.Println("SWARM SERVICES:           |")
			fmt.Println("--------------------------")
			listServices()
		case "3":
			fmt.Println("--------------------------")
			fmt.Println("SWARM NODES:              |")
			fmt.Println("--------------------------")
			listSwarmNodes()
		case "0":
			menuPage()
			return
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}

func initPage() {
	var selected string
	items := []string{
		"1. Docker stack init/update",
		"2. Remove previously deployed stack",
		"0. Back",
	}
	border := createBorder(items)
	renderMenu(items, border, "STACK MENU")

	for selected != "0" {
		fmt.Println(border)
		fmt.Print("Choose an option: ")
		selected = readInput()

		switch selected {
		case "1":
			stackInit()
		case "2":
			removeStack()
		case "0":
			menuPage()
			return
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}
