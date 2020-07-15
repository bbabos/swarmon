package main

import (
	"github.com/bbabos/swarmon/cmd/page"
	"github.com/bbabos/swarmon/cmd/stack"
	"github.com/bbabos/swarmon/cmd/utils"
	"github.com/bbabos/swarmon/config"
)

func main() {
	configExist := utils.FileExists(config.Paths.StackConfig)

	if configExist {
		config.Load(config.Paths.StackConfig)
		stack.SetAnswers()
	}
	page.MainPage()
}
