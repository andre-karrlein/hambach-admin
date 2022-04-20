package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type metrics struct {
	app.Compo
}

func (m *metrics) OnNav(ctx app.Context) {
	var status Status

	ctx.SessionStorage().Get("status", &status)

	if !status.LoggedIn {
		ctx.Navigate("/")
	}
}

func (m *metrics) Render() app.UI {
	return app.Div().Body(
		&navbar{},
		app.Section().Class("section").Body(
			app.Div().Class("has-navbar-fixed-top body").Body(
				app.Div().Class("container").Body(
					app.Div().Class("columns").Body(
						app.Div().Class("column is-one-fifth").Body(
							&sidebar{},
						),
						app.Div().Class("column").Body(
							app.P().Text("comming soon..."),
						),
					),
				),
			),
		),
	)
}
