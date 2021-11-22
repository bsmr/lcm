package page

import "github.com/bsmr/lcm/widget"

type Page struct {
	identifier string
	title      string
	widgets    []widget.Widget
}

func New(id, title string, widgets ...widget.Widget) (*Page, error) {
	page := &Page{
		identifier: id,
		title:      title,
		widgets:    widgets,
	}

	return page, nil
}

func (w *Page) Identifier() string {
	return w.identifier
}
