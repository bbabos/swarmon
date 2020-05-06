package page

import "fmt"

type page struct {
	title   string
	options string
	action  func()
}

var border string

func renderMenu(items []string, title string) {
	border = createMenuSeparator(items)
	clear()
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

func renderPage(pageitem page) {
	clear()
	separator := createOptionSeparator(pageitem.options)
	addSeparator(pageitem.title, separator)
	pageitem.action()
	addSeparator(pageitem.options, separator)
}

func addSeparator(header string, separator string) {
	fmt.Println(separator)
	fmt.Println(header)
	fmt.Println(separator)
}
