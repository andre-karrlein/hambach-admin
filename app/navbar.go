package main

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type navbar struct {
	app.Compo

	active string
}

func (n *navbar) OnClick(ctx app.Context, e app.Event) {
	status := Status{LoggedIn: false}
	ctx.SessionStorage().Set("status", status)
	ctx.Navigate("/")
}

func (n *navbar) Render() app.UI {
	return app.Nav().Class("navbar is-success is-fixed-top").Body(
		app.Div().Class("navbar-brand").Body(
			app.A().Class("navbar-item").Href("/").Body(
				app.Img().Src("https://storage.googleapis.com/hambach/IMG_0265.JPG"),
				app.H1().Class("title title-brand").Text("ADMIN"),
			),
			app.Span().Class("navbar-burger").Class(n.active).Body(
				app.Span(),
				app.Span(),
				app.Span(),
			).OnClick(n.onClick),
		),
		app.Div().Class("navbar-menu").ID("navbarMenu").Class(n.active).Body(
			app.Div().Class("navbar-start").Body(
				app.Div().Class("navbar-item").Body(),
			),
			app.Div().Class("navbar-end").Body(
				app.Div().Class("navbar-item").Body(
					app.Div().Class("buttons").Body(
						app.Button().Class("button is-danger").Text("Logout").OnClick(n.OnClick),
					),
				),
			),
		),
	)
}

func (n *navbar) onClick(ctx app.Context, e app.Event) {
	if n.active == "is-active" {
		n.active = ""
	} else {
		n.active = "is-active"
	}
	n.Update()
}
