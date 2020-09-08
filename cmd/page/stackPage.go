package page

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/bbabos/swarmon/cmd/config"
	"github.com/bbabos/swarmon/cmd/utils"
)

func stackInitOrUpdate() {
	var final string
	var msg string
	border := "----------------------------------------------"
	stackExist := stackExistCheck()

	if stackExist {
		msg = "Update existing monitoring stack..."
	} else {
		msg = "New monitoring stack initialization started..."
	}
	final = border + "\n" + msg + "\n" + border
	fmt.Println(final)

	config.GetAnswers(stackExist)
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
	if stackExist {
		msg = "Services updated succesfully..."
	} else {
		msg = "Stack deployed succesfully..."
	}
	final = border + "\n" + msg
	fmt.Println(final)
	utils.ExitOnKeystroke(stackPage)
}

func stackDelete() {
	stackExist := stackExistCheck()
	if stackExist {
		fmt.Print("Are you sure? [y/N]: ")
		input := utils.ReadInput()
		if input == "y" {
			utils.ExecShellCommand("docker stack rm "+config.Params.Docker.StackName, true)
			config.Params.Docker.StackName = ""
			config.CreateOrSave(config.Paths.StackConfig)
			fmt.Println("----------------------------------------------")
			fmt.Println("Monitoring stack deleted successfully!")
			utils.ExitOnKeystroke(stackPage)
		} else {
			stackPage()
		}
	} else {
		fmt.Println("You may not have a monitoring stack deployed!")
		utils.ExitOnKeystroke(stackPage)
	}
}

func stackExistCheck() bool {
	if config.Params.Docker.StackName == "" {
		return false
	}
	var out bytes.Buffer
	cmd := exec.Command("docker", "stack", "ls", "--format", "'{{.Name}}'")
	cmd.Stdout = &out
	cmd.Run()
	stdout := out.String()
	contains := strings.Contains(stdout, config.Params.Docker.StackName)
	return contains
}
