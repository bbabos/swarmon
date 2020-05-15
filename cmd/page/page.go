package page

import (
	"fmt"
)

type iPage interface {
	renderMenuPage()
	createBorder()
	renderSeparator()
}

type menuPage struct {
	title     string
	border    string
	menuItems []string
}

func (p *menuPage) renderMenuPage() {
	// utils.Clear()
	p.createBorder()
	p.renderSeparator()
	for _, item := range p.menuItems {
		fmt.Println(item)
	}
	fmt.Println(p.border)
}

func (p *menuPage) createBorder() {
	border, length := "", 0
	for _, item := range p.menuItems {
		if len(item) > length {
			length = len(item)
		}
	}
	for i := 0; i < length; i++ {
		border += "-"
	}
	p.border = border + border
}

func (p *menuPage) renderSeparator() {
	fmt.Println(p.border)
	fmt.Println(p.title)
	fmt.Println(p.border)
}
