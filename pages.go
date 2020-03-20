package main

import "fmt"

func menuPage() {
	// clear()
	fmt.Println("1. Deploy monitoring stack to swarm")
	fmt.Println("2. Maintain monitoring services")
	fmt.Println("3. Exit")

	validInput := true

	for validInput {
		fmt.Print("Choose an option: ")
		choosen := readInput()

		switch choosen {
		case "1":
			validInput = false
			stackInit()
		case "2":
			validInput = false
			fmt.Println("TODO")
		case "3":
			validInput = false
		default:
			fmt.Printf("%s is not a valid option\n", choosen)
		}
	}
}
