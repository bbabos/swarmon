package main

import (
	"github.com/bbabos/swarmon/cmd/config"
	"github.com/bbabos/swarmon/cmd/page"
	"github.com/bbabos/swarmon/cmd/utils"
)

func main() {
	configExists := utils.FileExists(config.Paths.StackConfig)
	if configExists {
		config.Load(config.Paths.StackConfig)
	}
	config.SetAnswers()
	page.MainPage()
}
