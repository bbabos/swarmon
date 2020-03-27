package main

import "fmt"

var border string

func renderPage(items []string, title string) {
	border = createBorder(items)
	clear()
	fmt.Println(border)
	fmt.Println(title)
	fmt.Println(border)
	for _, item := range items {
		fmt.Println(item)
	}
}

func createBorder(items []string) string {
	border, length := "", 0
	for _, item := range items {
		if len(item) > length {
			length = len(item)
		}
	}
	for i := 0; i < length; i++ {
		border += "-"
	}
	return border
}
