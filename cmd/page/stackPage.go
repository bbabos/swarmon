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
	stackExists := stackExistCheck()

	renderInfoText(stackExists, "Update existing monitoring stack...", "New monitoring stack initialization started...", false)
	config.GetAnswers(stackExists)
	parsedFile := utils.ParseFile(config.Paths.RawStack, config.Params)
	utils.WriteToFile(parsedFile, config.Paths.ParsedStack)

	renderInfoText(stackExists, "Updating docker services...", "Stack deploy started...", false)
	utils.ExecShellCommand("docker stack deploy -c "+config.Paths.ParsedStack+" "+config.Params.Docker.StackName, true)

	renderInfoText(stackExists, "Services updated succesfully...", "Stack deployed succesfully...", true)
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

func renderInfoText(stackExists bool, existMsg string, nonExistMsg string, lastText bool) {
	var msg string
	var final string
	border := "----------------------------------------------"

	if stackExists {
		msg = existMsg
	} else {
		msg = nonExistMsg
	}
	if lastText {
		final = border + "\n" + msg
	} else {
		final = border + "\n" + msg + "\n" + border
	}
	fmt.Println(final)

}
