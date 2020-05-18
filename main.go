package main

import (
	"github.com/bbabos/swarmon/cmd/page"
	"github.com/bbabos/swarmon/cmd/utils"
	"github.com/bbabos/swarmon/config"
)

func main() {
	configexist := utils.FileExists(config.ConfigPath)

	if configexist {
		utils.LoadConfig(config.ConfigPath)
		page.SetAnswers()
	}
	page.MainPage()
}
