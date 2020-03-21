package main

import "fmt"

type menuItem struct {
	option string
	action func()
}

func renderMenu(items []menuItem, border string, title string) {
	fmt.Println(border)
	fmt.Println(title)
	fmt.Println(border)
	for _, item := range items {
		fmt.Println(item.option)
	}
	fmt.Println(border)
}

func createBorder(items []menuItem) string {
	border := ""
	length := 0
	for _, item := range items {
		if len(item.option) > length {
			length = len(item.option)
		}
	}
	for i := 0; i < length; i++ {
		border += "-"
	}
	return border
}

func menuPage() {
	validInput := true
	items := []menuItem{
		{option: "1. Deploy monitoring stack to swarm", action: initPage},
		{option: "2. Maintain monitoring services", action: dockerPage},
		{option: "3. Exit"},
	}
	border := createBorder(items)
	renderMenu(items, border, "MAIN MENU")

	for validInput {
		fmt.Print("Choose an option: ")
		choosen := readInput()

		switch choosen {
		case "1":
			validInput = false
			items[0].action()
		case "2":
			validInput = false
			items[1].action()
		case "3":
			validInput = false
		default:
			fmt.Printf("%s is not a valid option\n", choosen)
		}
	}
}

func dockerPage() {
	validInput := true
	items := []menuItem{
		{option: "1. List all running container IDs", action: listContainerIDs},
		{option: "2. List services", action: listServices},
		{option: "3. List swarm nodes", action: listSwarmNodes},
		{option: "4. Exit to main menu", action: menuPage},
	}
	border := createBorder(items)
	renderMenu(items, border, "DOCKER MENU")

	for validInput {
		fmt.Print("Choose an option: ")
		choosen := readInput()

		switch choosen {
		case "1":
			items[0].action()
		case "2":
			items[1].action()
		case "3":
			items[2].action()
		case "4":
			validInput = false
			items[3].action()
		default:
			fmt.Printf("%s is not a valid option\n", choosen)
		}
	}
}

func initPage() {
	fmt.Println()
	fmt.Println("Swarm stack initialization started...")
	gitClone("https://github.com/babobene/swarmon.git", "tmp")
	getAnswers()
	parsedfile := parseFile("tmp/docker-compose.yml", p)
	writeToFile(parsedfile, "tmp/parsed.yml")
	stackDeploy("tmp/parsed.yml", p.Docker.StackName)
}