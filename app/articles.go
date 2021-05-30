package main

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type articles struct {
	app.Compo
}

func (a *articles) OnNav(ctx app.Context) {
	var status Status

	ctx.SessionStorage().Get("status", &status)

	if !status.LoggedIn {
		ctx.Navigate("/")
	}
}

func (a *articles) Render() app.UI {
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
							&articleList{},
						),
					),
				),
			),
		),
	)
}
