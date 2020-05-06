package main

import (
	"github.com/bbabos/swarmon-go/cmd/page"
	"github.com/bbabos/swarmon-go/cmd/utils"
	"github.com/bbabos/swarmon-go/config"
)

func main() {
	configexist := utils.FileExists(config.ConfigPath)

	if configexist {
		utils.LoadConfig(config.ConfigPath)
		page.SetAnswers()
	}
	page.MenuPage()
}
