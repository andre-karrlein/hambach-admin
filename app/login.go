package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"golang.org/x/crypto/bcrypt"
)

type login struct {
	app.Compo
}

func (l *login) OnNav(ctx app.Context) {
	var status Status

	ctx.SessionStorage().Get("status", &status)

	if status.LoggedIn {
		ctx.Navigate("/home")
	}
}

func (l *login) Render() app.UI {
	return app.Section().Class("section").Body(
		app.Div().Body(
			app.Div().Class("container").Body(
				app.Form().Class("box").OnSubmit(l.OnSubmit).Body(
					app.Div().Class("field").Body(
						app.Label().Class("label").Text("User"),
						app.Div().Class("control").Body(
							app.Input().Class("input").Type("text").Placeholder("username").ID("username"),
						),
					),
					app.Div().Class("field").Body(
						app.Label().Class("label").Text("Password"),
						app.Div().Class("control").Body(
							app.Input().Class("input").Type("password").Placeholder("********").ID("password"),
						),
					),
					app.Button().Class("button is-primary").Text("Sign in").Type("submit"),
				),
			),
		),
	)
}

func (l *login) OnSubmit(ctx app.Context, e app.Event) {
	e.PreventDefault()
	username := app.Window().GetElementByID("username").Get("value").String()
	password := app.Window().GetElementByID("password").Get("value").String()

	if loginUser(username, password) {
		status := Status{LoggedIn: true, User: username}
		ctx.SessionStorage().Set("status", status)
		ctx.Navigate("/home")
	} else {
		status := Status{LoggedIn: false, User: ""}
		ctx.SessionStorage().Set("status", status)
	}
}

func loginUser(username, password string) bool {
	users := GetUsers()
	if val, ok := users[username]; ok {
		return checkPasswordHash(password, val.password)
	}

	return false
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
