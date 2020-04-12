package main

import "fmt"

func serviceOptions() {
	var selected string
	pageItem := page{
		title:   "SERVICES",
		options: "| 0 - Back | 1 - another opts | 2 - another opts |",
		action:  listServices,
	}
	renderPage(pageItem)

loop:
	for {
		fmt.Print("Select an option: ")
		selected = readInput()

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
