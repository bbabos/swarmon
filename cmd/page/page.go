package page

import (
	"fmt"

	"github.com/bbabos/swarmon/cmd/utils"
)

type page struct {
	title     string
	border    string
	menuItems []string
	options   string
	action    func()
}

func renderMenuPage(p *page) {
	utils.Clear()
	p.createBorder()
	p.renderSeparator()
	for _, item := range p.menuItems {
		fmt.Println(item)
	}
	fmt.Println(p.border)
}

func renderSubPage(p *page) {
	utils.Clear()
	p.createOptionBorder()
	p.renderSeparator()
	p.action()
	fmt.Println(p.border)
	fmt.Println(p.options)
	fmt.Println(p.border)
}

func (p *page) createBorder() {
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

func (p *page) renderSeparator() {
	fmt.Println(p.border)
	fmt.Println(p.title)
	fmt.Println(p.border)
}

func (p *page) createOptionBorder() {
	length := len(p.options)
	border := ""

	for i := 0; i < length; i++ {
		border += "-"
	}
	p.border = border
}
