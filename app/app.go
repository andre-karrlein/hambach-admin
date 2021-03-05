package main

import (
	"log"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"golang.org/x/crypto/bcrypt"
)

type navbar struct {
	app.Compo
}

type status struct {
	LoggedIn bool `json:"loggedIn"`
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

	Username string
	Password string
}

const hash = "$2y$14$7aNuDEs7G6KxyYZLShEHlOpY4cjxV4kizm3noGFNBW11dvJdgtp3G"

func getUsers() map[string]string {
	return map[string]string{
		"akarrlein": hash,
		"pgeissler": hash,
	}
}

func (h *home) Render() app.UI {
	var status status
	app.SessionStorage.Get("status", &status)

	if status.LoggedIn != true {
		app.Navigate("/login")
	}
	return app.Div().Body(
		app.Button().Class("button is-danger").Text("Logout").OnClick(h.OnClick),
	)
}

func (h *home) OnClick(ctx app.Context, e app.Event) {
	status := status{LoggedIn: false}
	app.SessionStorage.Set("status", status)
	app.Navigate("/login")
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
		status := status{LoggedIn: true}
		app.SessionStorage.Set("status", status)
		app.Navigate("/")
	} else {
		status := status{LoggedIn: false}
		app.SessionStorage.Set("status", status)
	}
}

func loginUser(username, password string) bool {
	log.Println(username)
	log.Println(password)
	users := getUsers()
	if val, ok := users[username]; ok {
		log.Println(val)
		return checkPasswordHash(password, val)
	}

	return false
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	app.Route("/", &home{})
	app.Route("/login", &login{})
	app.Route("/article", &article{})
	app.Run()
}
