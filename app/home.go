package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type home struct {
	app.Compo

	status Status
}

// Status reflects the logged in status
type Status struct {
	LoggedIn bool   `json:"loggedIn"`
	User     string `json:"user"`
}

func (h *home) OnNav(ctx app.Context) {
	var status Status

	ctx.SessionStorage().Get("status", &status)

	if !status.LoggedIn {
		ctx.Navigate("/")
	}

	h.status = status
	h.Update()
}

func (h *home) Render() app.UI {
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
							app.H1().Text("Hallo "+GetUsers()[h.status.User].name),
							app.P().Text("comming soon test..."),
						),
					),
				),
			),
		),
	)
}
