package main

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type editor struct {
	app.Compo
}

func (e *editor) Render() app.UI {
	return app.Div().Body(
		&navbar{},
		app.Section().Class("section").Body(
			app.Body().Class("has-navbar-fixed-top").Body(
				app.Div().Class("container").Body(
					app.Div().Class("columns").Body(
						app.Div().Class("column is-one-fifth").Body(
							&sidebar{},
						),
						app.Div().Class("column").Body(
						//&article{},
						),
					),
				),
			),
		),
	)
}
