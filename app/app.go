package main

import (
	"log"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type navbar struct {
	app.Compo
}

type home struct {
	app.Compo

	loggedIn bool
}

type article struct {
	app.Compo

	loggedIn bool
}

type login struct {
	app.Compo

	username string
	password string
}

func (h *home) Render() app.UI {
	var status string
	app.SessionStorage.Get("loggedIn", status)
	//if status == "" {
	//	status = "false"
	//}
	if status == "false" {
		//~app.Navigate("/login")
	}
	return app.Div().Body(
		app.Button().Class("button is-danger").Text(status).OnClick(h.OnClick),
	)
}

func (h *home) OnClick(ctx app.Context, e app.Event) {
	//app.SessionStorage.Set("loggedIn", "false")
	app.Navigate("/login")
}

func (l *login) Render() app.UI {
	return app.Div().Body(
		//&navbar{},
		app.Form().Class("box").Body(
			app.Div().Class("field").Body(
				app.Label().Class("label").Text("User"),
				app.Div().Class("control").Body(
					app.Input().Class("input").Type("text").Placeholder("username").Value(l.username),
				),
			),
			app.Div().Class("field").Body(
				app.Label().Class("label").Text("Password"),
				app.Div().Class("control").Body(
					app.Input().Class("input").Type("password").Placeholder("********").Value(l.password),
				),
			),
			app.Button().Class("button is-primary").Text("Sign in").OnClick(l.OnClick),
		),
	)
}

func (l *login) OnClick(ctx app.Context, e app.Event) {
	//ctx.JSSrc.Get("value")
	status := "True"
	log.Fatal(app.SessionStorage.Set("loggedIn", status))
	app.Navigate("/")
}

func main() {
	app.Route("/", &home{})
	app.Route("/login", &login{})
	app.Route("/article", &article{})
	app.Run()
}
