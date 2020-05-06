package main

import (
	"github.com/bbabos/swarmon-go/cmd/page/menupage"
	"github.com/bbabos/swarmon-go/cmd/page/stackpage"
	"github.com/bbabos/swarmon-go/cmd/utils"
	"github.com/bbabos/swarmon-go/config"
)

func main() {
	configexist := utils.FileExists(config.ConfigPath)

	if configexist {
		utils.LoadConfig(config.ConfigPath)
		stackpage.SetAnswers()
	}
	menupage.MenuPage()
}
