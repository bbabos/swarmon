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
			stackInitPage()
		case "2":
			validInput = false
			dockerPage()
		case "3":
			validInput = false
		default:
			fmt.Printf("%s is not a valid option\n", choosen)
		}
	}
}

func stackInitPage() {
	fmt.Println()
	fmt.Println("Swarm stack initialization started...")
	// gitClone("https://github.com/babobene/swarmon.git", "tmp")
	getAnswers()
	parsedfile := parseFile("templates/example.yml", p)
	writeToFile(parsedfile, "templates/parsed.yml")
	// stackDeploy("templates/parsed.yml", p.Docker.StackName)
}

func dockerPage() {
	fmt.Println("1. List all running container IDs")
	fmt.Println("2. List services")
	fmt.Println("3. List swarm nodes")
	fmt.Println("4. Exit to main menu")

	validInput := true

	for validInput {
		fmt.Print("Choose an option: ")
		choosen := readInput()

		switch choosen {
		case "1":
			listContainerIDs()
			fmt.Println()
		case "2":
			listServices()
			fmt.Println()
		case "3":
			listSwarmNodes()
			fmt.Println()
		case "4":
			validInput = false
			menuPage()
		default:
			fmt.Printf("%s is not a valid option\n", choosen)
		}
	}
}
