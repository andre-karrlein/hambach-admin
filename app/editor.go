package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type editor struct {
	app.Compo

	item Content
	file File
}

func (e *editor) OnNav(ctx app.Context) {
	var status Status

	ctx.SessionStorage().Get("status", &status)

	if !status.LoggedIn {
		ctx.Navigate("/")
	}

	urlPath := ctx.Page().URL().Path
	path := strings.Split(urlPath, "/")
	id := path[2]

	if path[3] == "new" {
		e.getDefaultItem(id)
		return
	}

	// Launching a new goroutine:
	ctx.Async(func() {
		app_key := app.Getenv("READ_KEY")
		r, err := http.Get("https://api.spvgg-hambach.de/api/v1/content/" + id + "?appkey=" + app_key)
		if err != nil {
			app.Log(err)
			return
		}
		defer r.Body.Close()

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			app.Log(err)
			return
		}

		sb := string(b)

		var content Content
		json.Unmarshal([]byte(sb), &content)

		e.item = content
		e.Update()
	})
}

func (e *editor) OnMount(ctx app.Context) {
	ctx.ObserveState("file").Value(&e.file)
	log.Println(e.file.Key)
	e.Update()
}

func (e *editor) onClick(ctx app.Context, ev app.Event) {
	ctx.Navigate("/articles")
}

func (e *editor) getDefaultItem(articleID string) {
	e.item = Content{
		ID:      articleID,
		Image:   "https://hambach.s3.eu-central-1.amazonaws.com/hambach_logo.png",
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
	app_key := app.Getenv("WRITE_KEY")
	id := app.Window().GetElementByID("id").Get("textContent").String()
	title := app.Window().GetElementByID("title").Get("value").String()
	link := app.Window().GetElementByID("link").Get("value").String()
	image := app.Window().GetElementByID("image").Get("value").String()
	category := app.Window().GetElementByID("category").Get("value").String()
	date := app.Window().GetElementByID("date").Get("value").String()
	creator := app.Window().GetElementByID("creator").Get("value").String()
	content := app.Window().GetElementByID("content").Get("value").String()
	active := app.Window().GetElementByID("active").Get("value").String()

	article := Content{
		ID:       id,
		Sort_key: date,
		Title:    title,
		Date:     date,
		Category: category,
		Type:     contentType,
		Image:    image,
		Creator:  creator,
		Content:  content,
		Active:   active,
		Link:     link,
	}

	// marshal User to json
	json, err := json.Marshal(article)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	// set the HTTP method, url, and request body
	req, err := http.NewRequest(http.MethodPut, "https://api.spvgg-hambach.de/api/v1/content/?appkey="+app_key, bytes.NewBuffer(json))
	if err != nil {
		panic(err)
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}

func (e *editor) Render() app.UI {
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
