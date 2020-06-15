package main

import (
	"github.com/bbabos/swarmon/cmd/page"
	"github.com/bbabos/swarmon/cmd/utils"
	"github.com/bbabos/swarmon/config"
)

func main() {
	configExist := utils.FileExists(config.Path)

	if configExist {
		config.Load(config.Path)
		page.SetAnswers()
	}
	page.MainPage()
}
