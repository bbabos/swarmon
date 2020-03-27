package main

import "fmt"

func containerOptions() {
	var selected string
	exit := false

	clear()
	fmt.Println(border)
	fmt.Println("CONTAINERS")
	fmt.Println(border)
	listContainers()

	for !exit {
		fmt.Println(border)
		fmt.Println("0 - Back | 1 - Enter into a container")
		fmt.Println(border)
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
