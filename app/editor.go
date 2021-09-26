package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type editor struct {
	app.Compo

	item Content
}

func (e *editor) OnNav(ctx app.Context) {
	var status Status

	ctx.SessionStorage().Get("status", &status)

	if !status.LoggedIn {
		ctx.Navigate("/")
	}

	urlPath := ctx.Page.URL().Path
	path := strings.Split(urlPath, "/")
	id := path[2]

	if path[3] == "new" {
		e.getDefaultItem(id)
		return
	}

	// Launching a new goroutine:
	ctx.Async(func() {
		r, err := http.Get("/api/v1/content?id=" + id) // #TODO
		if err != nil {
			app.Log(err)
			return
		}
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			app.Log(err)
			return
		}

		// Storing HTTP response in component field:
		ctx.Dispatch(func() {
			sb := string(body)

			var content Content
			json.Unmarshal([]byte(sb), &content)

			e.item = content
			e.Update()
		})
	})
}

func (e *editor) onClick(ctx app.Context, ev app.Event) {
	ctx.Navigate("/articles")
}

func (e *editor) getDefaultItem(articleID string) {
	contentID, _ := strconv.Atoi(articleID)
	e.item = Content{
		ID:      contentID,
		Image:   "https://storage.googleapis.com/hambach/hambach_logo.png",
		Content: "<div class=\"matches\"><p></p></div>",
	}
	e.Update()
}

func (e *editor) OnSubmit(ctx app.Context, ev app.Event) {
	ev.PreventDefault()
	content := app.Window().GetElementByID("type").Get("value").String()
	e.save(content)

	ctx.Navigate("/articles")
}

func (e *editor) OnDeactivate(ctx app.Context, ev app.Event) {
	ev.PreventDefault()
	e.save("content")

	ctx.Navigate("/articles")
}

func (e *editor) OnActivate(ctx app.Context, ev app.Event) {
	ev.PreventDefault()
	e.save("article")

	ctx.Navigate("/articles")
}

func (e *editor) save(contentType string) {
	id := app.Window().GetElementByID("id").Get("textContent").String()
	title := app.Window().GetElementByID("title").Get("value").String()
	link := app.Window().GetElementByID("link").Get("value").String()
	image := app.Window().GetElementByID("image").Get("value").String()
	category := app.Window().GetElementByID("category").Get("value").String()
	date := app.Window().GetElementByID("date").Get("value").String()
	creator := app.Window().GetElementByID("creator").Get("value").String()
	content := app.Window().GetElementByID("content").Get("value").String()
	active := app.Window().GetElementByID("active").Get("value").String()

	data := url.Values{
		"id":          {id},
		"title":       {title},
		"date":        {date},
		"category":    {category},
		"contentType": {contentType},
		"image":       {image},
		"creator":     {creator},
		"content":     {content},
		"active":      {active},
		"link":        {link},
	}

	_, err := http.PostForm("/api/v1/content/save", data)
	if err != nil {
		log.Fatal(err)
	}
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
							app.Div().Body(
								app.Div().Class("buttons").Body(
									app.Button().Class("button").Text("Zurück").OnClick(e.onClick),
									app.Button().Class("button is-danger").Text("Deaktivieren").OnClick(e.OnDeactivate),
									app.Button().Class("button is-success").Text("Aktivieren").OnClick(e.OnActivate),
								),
								app.Br(),
								app.Br(),
								app.Form().Class("box").OnSubmit(e.OnSubmit).Body(
									app.Label().Text(e.item.ID).ID("id"),
									app.Br(),
									app.Br(),
									app.Div().Class("field").Body(
										app.Label().Class("label").Text("Titel"),
										app.Div().Class("control").Body(
											app.Input().Class("input").Type("text").Value(e.item.Title).ID("title"),
										),
									),
									app.Div().Class("field").Body(
										app.Label().Class("label").Text("Link im Titel?"),
										app.Div().Class("control").Body(
											app.Input().Class("input").Type("text").Value(e.item.Link).ID("link"),
										),
									),
									app.Div().Class("field").Body(
										app.Label().Class("label").Text("Titelbild"),
										app.Div().Class("control").Body(
											app.Input().Class("input").Type("text").Value(e.item.Image).ID("image"),
											//https://storage.googleapis.com/hambach/hambach_logo.png
										),
									),
									app.Div().Class("field").Body(
										app.Label().Class("label").Text("Kategorie"),
										app.Div().Class("control").Body(
											app.Input().Class("input").Type("text").Value(e.item.Category).ID("category"),
										),
									),
									app.Div().Class("field").Body(
										app.Label().Class("label").Text("Datum"),
										app.Div().Class("control").Body(
											app.Input().Class("input").Type("date").Value(e.item.Date).ID("date"),
										),
									),
									app.Div().Class("field").Body(
										app.Label().Class("label").Text("Ersteller"),
										app.Div().Class("control").Body(
											app.Input().Class("input").Type("text").Value(e.item.Creator).ID("creator"),
										),
									),
									app.Div().Class("field").Body(
										app.Label().Class("label").Text("Typ"),
										app.Div().Class("control").Body(
											app.Input().Class("input").Type("text").Value(e.item.Type).ID("type"),
										),
									),
									app.Div().Class("field").Body(
										app.Label().Class("label").Text("Inhalt"),
										app.Div().Class("control").Body(
											app.Textarea().Class("textarea").Text(e.item.Content).ID("content"),
										),
									),
									app.Div().Class("field").Body(
										app.Label().Class("label").Text("Aktiviert"),
										app.Div().Class("control").Body(
											app.Input().Class("input").Type("text").Value(e.item.Active).ID("active"),
										),
									),
									app.Button().Class("button is-primary").Text("Speichern").Type("submit"),
								),
							),
						),
					),
				),
			),
		),
	)
}
