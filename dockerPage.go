package main

import "fmt"

func serviceOptions() {
	var selected string
	exit := false
	pageItem := page{
		title:   "SERVICES",
		options: "| 0 - Back | 1 - another opts | 2 - another opts |",
		action:  listServices,
	}
	renderPage(pageItem)

	for !exit {
		fmt.Print("Select an option: ")
		selected = readInput()

		switch selected {
		case "1":
			// TODO
		case "2":
			// TODO
		case "0":
			exit = true
			dockerPage()
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}
