package main

import (
	"github.com/bbabos/swarmon/cmd/config"
	"github.com/bbabos/swarmon/cmd/page"
	"github.com/bbabos/swarmon/cmd/utils"
)

func main() {
	configExist := utils.FileExists(config.Paths.StackConfig)

	if configExist {
		config.Load(config.Paths.StackConfig)
		config.SetAnswers()
	}
	page.MainPage()
}
