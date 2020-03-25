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
	// fmt.Println(border)
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
	validInput := true
	items := []string{
		"1. Monitoring stack options",
		"2. Maintain monitoring services",
		"0. Exit",
	}
	border := createBorder(items)
	renderMenu(items, border, "MAIN MENU")

	for validInput {
		fmt.Println(border)
		fmt.Print("Choose an option: ")
		choosen := readInput()

		switch choosen {
		case "1":
			validInput = false
			initPage()
		case "2":
			validInput = false
			dockerPage()
		case "0":
			validInput = false
		default:
			fmt.Printf("%s is not a valid option\n", choosen)
		}
	}
}

func dockerPage() {
	validInput := true
	items := []string{
		"1. List all running containers",
		"2. List services",
		"3. List swarm nodes",
		"0. Exit to main menu",
	}
	border := createBorder(items)
	renderMenu(items, border, "DOCKER MENU")

	for validInput {
		fmt.Println(border)
		fmt.Print("Choose an option: ")
		choosen := readInput()

		switch choosen {
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
			validInput = false
			menuPage()
		default:
			fmt.Printf("%s is not a valid option\n", choosen)
		}
	}
}

func initPage() {
	validInput := true
	items := []string{
		"1. Docker stack init/update",
		"2. Remove previously deployed stack",
		"0. Exit to main menu",
	}
	border := createBorder(items)
	renderMenu(items, border, "STACK MENU")

	for validInput {
		fmt.Println(border)
		fmt.Print("Choose an option: ")
		choosen := readInput()

		switch choosen {
		case "1":
			stackInit()
		case "2":
			removeStack()
		case "0":
			validInput = false
			menuPage()
		default:
			fmt.Printf("%s is not a valid option\n", choosen)
		}
	}
}
