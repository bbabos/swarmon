package page

import "fmt"

type iSubPage interface {
	renderSubPage()
	createOptionBorder()
	renderSeparator()
}

type subPage struct {
	action  func()
	options string
	border  string
	title   string
}

func (p *subPage) renderSubPage() {
	// utils.Clear()
	p.createOptionBorder()
	p.renderSeparator()
	p.action()
	p.renderSeparator()
}

func (p *subPage) createOptionBorder() {
	length := len(p.options)
	border := ""

	for i := 0; i < length; i++ {
		border += "-"
	}
	p.border = border
}

func (p *subPage) renderSeparator() {
	fmt.Println(p.border)
	fmt.Println(p.title)
	fmt.Println(p.border)
}
