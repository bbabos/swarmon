package page

import (
	"fmt"

	"github.com/bbabos/swarmon-go/cmd/utils"
)

// Page is ...
type Page struct {
	Title   string
	Options string
	Action  func()
}

var border string

func renderMenu(items []string, title string) {
	border = createMenuSeparator(items)
	utils.Clear()
	fmt.Println(border)
	fmt.Println(title)
	fmt.Println(border)
	for _, item := range items {
		fmt.Println(item)
	}
	fmt.Println(border)
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

func renderPage(pageitem Page) {
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
