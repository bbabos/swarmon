package main

import (
	"github.com/bbabos/swarmon/cmd/page"
	"github.com/bbabos/swarmon/cmd/utils"
	"github.com/bbabos/swarmon/config"
)

func main() {
	configexist := utils.FileExists(config.Path)

	if configexist {
		utils.LoadConfig(config.Path)
		page.SetAnswers()
	}
	page.MainPage()
}
