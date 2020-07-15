package stack

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/bbabos/swarmon/config"
)

// SetAnswers is ...
func SetAnswers() {
	config.Inputs[0].Answer = config.Params.Docker.StackName
	config.Inputs[1].Answer = config.Params.Domain
	config.Inputs[2].Answer = config.Params.AdminUser.Name
	config.Inputs[3].Answer = config.Params.AdminUser.Password
	config.Inputs[4].Answer = config.Params.Traefik.BAUser
	config.Inputs[5].Answer = config.Params.Traefik.BAPassword
	config.Inputs[6].Answer = config.Params.Slack.Webhook
	config.Inputs[7].Answer = config.Params.Slack.AlertUser
	config.Inputs[8].Answer = config.Params.Traefik.Port
	config.Inputs[9].Answer = config.Params.Schema
	config.Inputs[10].Answer = config.Params.Docker.MetricPort
	config.Inputs[11].Answer = config.Params.Docker.GwBridgeIP
}

// SetParams is ...
func SetParams() {
	config.Params.Docker.StackName = config.Inputs[0].Answer
	config.Params.Domain = config.Inputs[1].Answer
	config.Params.AdminUser.Name = config.Inputs[2].Answer
	config.Params.AdminUser.Password = config.Inputs[3].Answer
	config.Params.Traefik.BAUser = config.Inputs[4].Answer
	config.Params.Traefik.BAPassword = config.Inputs[5].Answer
	config.Params.Slack.Webhook = config.Inputs[6].Answer
	config.Params.Slack.AlertUser = config.Inputs[7].Answer
	config.Params.Traefik.Port = config.Inputs[8].Answer
	config.Params.Schema = config.Inputs[9].Answer
	config.Params.Docker.MetricPort = config.Inputs[10].Answer
	config.Params.Docker.GwBridgeIP = config.Inputs[11].Answer
}

// ExistCheck is ...
func ExistCheck() bool {
	var out bytes.Buffer
	cmd := exec.Command("docker", "stack", "ls", "--format", "'{{.Name}}'")
	cmd.Stdout = &out
	cmd.Run()
	stdout := out.String()
	contains := strings.Contains(stdout, config.Params.Docker.StackName)
	return contains
}
