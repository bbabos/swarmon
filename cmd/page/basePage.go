package page

import (
	"fmt"

	"github.com/bbabos/swarmon-go/cmd/utils"
	"github.com/bbabos/swarmon-go/config"
)

// Page is ...
type Page struct {
	Title   string
	Options string
	Action  func()
}

// RenderMenu is ...
func RenderMenu(items []string, title string) {
	config.Border = createMenuSeparator(items)
	utils.Clear()
	fmt.Println(config.Border)
	fmt.Println(title)
	fmt.Println(config.Border)
	for _, item := range items {
		fmt.Println(item)
	}
	fmt.Println(config.Border)
}

func createMenuSeparator(items []string) string {
	border, length := "", 0
	for _, item := range items {
		if len(item) > length {
			length = len(item)
		}
	}
	for i := 0; i < length; i++ {
		border += "-"
	}
	return border + border
}

func createOptionSeparator(options string) string {
	length := len(options)
	border := ""

	for i := 0; i < length; i++ {
		border += "-"
	}
	return border
}

// RenderPage is ...
func RenderPage(pageitem Page) {
	utils.Clear()
	separator := createOptionSeparator(pageitem.Options)
	addSeparator(pageitem.Title, separator)
	pageitem.Action()
	addSeparator(pageitem.Options, separator)
}

func addSeparator(header string, separator string) {
	fmt.Println(separator)
	fmt.Println(header)
	fmt.Println(separator)
}
