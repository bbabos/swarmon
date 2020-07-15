package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/stack"
	"github.com/bbabos/swarmon/cmd/utils"
	"github.com/bbabos/swarmon/config"
)

func getAnswers(stackExists bool) {
	length := len(config.Inputs)
	num := 0
	if stackExists {
		num = 1
	}
	for i := num; i < length; i++ {
		if config.Inputs[i].Answer == "" {
			fmt.Print(config.Inputs[i].Question + ": ")
			config.Inputs[i].Answer = utils.ReadInput()
		} else {
			fmt.Print(config.Inputs[i].Question + " [" + config.Inputs[i].Answer + "]" + ": ")
			result := utils.ReadInput()
			if result != "" {
				config.Inputs[i].Answer = result
			}
		}
	}
	stack.SetParams()
	config.Save(config.Paths.StackConfig)
	config.Params.Traefik.BAPassword = utils.HashPass(config.Inputs[5].Answer)
}

func stackInitOrUpdate() {
	var final string
	var msg string
	border := "----------------------------------------------"
	stackExist := stack.ExistCheck()

	if stackExist {
		msg = "Update existing monitoring stack..."
	} else {
		msg = "New monitoring stack initialization started..."
	}
	final = border + "\n" + msg + "\n" + border
	fmt.Println(final)

	getAnswers(stackExist)
	parsedFile := utils.ParseFile(config.Paths.RawStack, config.Params)
	utils.WriteToFile(parsedFile, config.Paths.ParsedStack)

	if stackExist {
		msg = "Updating docker services..."
	} else {
		msg = "Stack deploy started..."
	}
	final = border + "\n" + msg + "\n" + border
	fmt.Println(final)

	utils.ExecShellCommand("docker stack deploy -c "+config.Paths.ParsedStack+" "+config.Params.Docker.StackName, true)
	utils.ExitOnKeyStroke(stackPage)
}

func stackDelete() {
	stackExist := stack.ExistCheck()
	if stackExist {
		fmt.Print("Are you sure? [y/N]: ")
		input := utils.ReadInput()
		if input == "y" {
			utils.ExecShellCommand("docker stack rm "+config.Params.Docker.StackName, true)
			fmt.Println("-------------------------------")
			fmt.Println("Monitoring stack deleted successfully!")
		}
	} else {
		fmt.Println("You may not have a monitoring stack deployed!")
	}
	utils.ExitOnKeyStroke(stackPage)
}
