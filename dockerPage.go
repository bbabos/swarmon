package main

import "fmt"

func containerOptions() {
	var selected string
	exit := false
	pageItem := page{
		title:   "CONTAINERS",
		options: "| 0 - Back | 1 - Enter into a container |",
		action:  listContainers,
	}
	renderPage(pageItem)

	for !exit {
		fmt.Print("Select an option: ")
		selected = readInput()

		switch selected {
		case "1":
			// TODO
		case "0":
			exit = true
			dockerPage()
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}
