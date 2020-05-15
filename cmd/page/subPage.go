package page

import "fmt"

type iSubPage interface {
	renderSubPage()
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
	p.renderOptions()
}

func (p *subPage) createOptionBorder() {
	length := len(p.options)
	var border string

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

func (p *subPage) renderOptions() {
	fmt.Println(p.border)
	fmt.Println(p.options)
	fmt.Println(p.border)
}
