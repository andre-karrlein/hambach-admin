package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type sidebar struct {
	app.Compo
}

func (s *sidebar) Render() app.UI {
	return app.Aside().Class("menu").Body(
		app.P().Class("menu-label").Text("Administration"),
		app.Ul().Class("menu-list").Body(
			app.Li().Body(
				app.A().Href("https://eu-central-1.admin.amplifyapp.com/admin/d2gdxd5wqafh8p/staging/content").Text("Admin Neu"),
				app.A().Href("/").Text("Dashboard"),
				app.A().Href("/articles").Text("Artikel"),
				app.A().Href("/metrics").Text("Metriken"),
				app.A().Href("/files").Text("Dateien"),
			),
		),
	)
}
