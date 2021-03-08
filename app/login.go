package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"golang.org/x/crypto/bcrypt"
)

type login struct {
	app.Compo
}

func (l *login) Render() app.UI {
	return app.Section().Class("section").Body(
		app.Body().Body(
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
		status := Status{LoggedIn: true}
		app.SessionStorage.Set("status", status)
		app.Navigate("/")
	} else {
		status := Status{LoggedIn: false}
		app.SessionStorage.Set("status", status)
	}
}

func loginUser(username, password string) bool {
	users := GetUsers()
	if val, ok := users[username]; ok {
		return checkPasswordHash(password, val)
	}

	return false
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
