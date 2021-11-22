package application

import (
	"errors"

	"github.com/bsmr/lcm/page"
)

type Application struct {
	Pages []page.Page
}

func New(pages ...page.Page) (*Application, error) {

	app := &Application{
		Pages: pages,
	}

	return app, nil
}

func (a *Application) Run() error {
	return errors.New("not implemented")
}
