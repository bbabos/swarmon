package dockerpage

import (
	"fmt"

	"github.com/bbabos/swarmon-go/cmd/docker"
	"github.com/bbabos/swarmon-go/cmd/page"
	"github.com/bbabos/swarmon-go/cmd/page/menupage"
	"github.com/bbabos/swarmon-go/cmd/utils"
)

// ServiceOptions is ...
func ServiceOptions() {
	var selected string
	pageItem := page.Page{
		Title:   "SERVICES",
		Options: "| 0 - Back | 1 - another opts | 2 - another opts |",
		Action:  docker.ListServices,
	}
	page.RenderPage(pageItem)

loop:
	for {
		fmt.Print("Select an option: ")
		selected = utils.ReadInput()

		switch selected {
		case "1":
			// TODO
		case "2":
			// TODO
		case "0":
			menupage.DockerPage()
			break loop
		default:
			fmt.Printf("%s is not a valid option\n", selected)
		}
	}
}
