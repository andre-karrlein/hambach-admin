package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type article struct {
	app.Compo

	content   []Content
	edit      bool
	loading   bool
	articleID string
	item      Content
	nextID    string
}

// Content struct
type Content struct {
	ID       int
	Title    string
	Date     string
	Category string
	Type     string
	Image    string
	Creator  string
	Content  string
	Active   string
	Link     string
}

func (a *articles) Render() app.UI {
	var status Status
	app.SessionStorage.Get("status", &status)

	if !status.LoggedIn {
		app.Navigate("/login")
	}

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
							&article{},
						),
					),
				),
			),
		),
	)
}

func (a *article) Render() app.UI {
	if a.edit {

		return app.Div().Body(
			app.Button().Class("button").Text("Zurück").OnClick(a.onClick),
			app.Br(),
			app.Br(),
			app.Form().Class("box").OnSubmit(a.OnSubmit).Body(
				app.Label().Text(a.item.ID).ID("id"),
				app.Br(),
				app.Br(),
				app.Div().Class("field").Body(
					app.Label().Class("label").Text("Titel"),
					app.Div().Class("control").Body(
						app.Input().Class("input").Type("text").Value(a.item.Title).ID("title"),
					),
				),
				app.Div().Class("field").Body(
					app.Label().Class("label").Text("Link im Titel?"),
					app.Div().Class("control").Body(
						app.Input().Class("input").Type("text").Value(a.item.Link).ID("link"),
					),
				),
				app.Div().Class("field").Body(
					app.Label().Class("label").Text("Titelbild"),
					app.Div().Class("control").Body(
						app.Input().Class("input").Type("text").Value(a.item.Image).ID("image"),
						//https://storage.googleapis.com/hambach/hambach_logo.png
					),
				),
				app.Div().Class("field").Body(
					app.Label().Class("label").Text("Kategorie"),
					app.Div().Class("control").Body(
						app.Input().Class("input").Type("text").Value(a.item.Category).ID("category"),
					),
				),
				app.Div().Class("field").Body(
					app.Label().Class("label").Text("Datum"),
					app.Div().Class("control").Body(
						app.Input().Class("input").Type("date").Value(a.item.Date).ID("date"),
					),
				),
				app.Div().Class("field").Body(
					app.Label().Class("label").Text("Ersteller"),
					app.Div().Class("control").Body(
						app.Input().Class("input").Type("text").Value(a.item.Creator).ID("creator"),
					),
				),
				app.Div().Class("field").Body(
					app.Label().Class("label").Text("Typ"),
					app.Div().Class("control").Body(
						app.Input().Class("input").Type("text").Value(a.item.Type).ID("type"),
					),
				),
				app.Div().Class("field").Body(
					app.Label().Class("label").Text("Inhalt"),
					app.Div().Class("control").Body(
						app.Textarea().Class("textarea").Text(a.item.Content).ID("content"),
					),
				),
				app.Div().Class("field").Body(
					app.Label().Class("label").Text("Aktiviert"),
					app.Div().Class("control").Body(
						app.Input().Class("input").Type("text").Value(a.item.Active).ID("active"),
					),
				),
				app.Button().Class("button is-primary").Text("Speichern").Type("submit"),
			),
		)
	}

	return app.Div().Body(
		app.Button().Class("button is-success").ID(a.nextID).Text("Neuer Artikel").OnClick(a.newArticle),
		app.Br(),
		app.Br(),
		app.Range(a.content).Slice(func(i int) app.UI {
			color := ""
			if a.content[i].Type == "article" {
				color = "lightgreen"
			}

			return app.Div().Class("box").Style("background-color", color).Body(
				app.Article().Class("media").Body(
					app.Div().Class("media-left").Body(
						app.Figure().Class("image is-64x64").Body(
							app.Img().Src(a.content[i].Image),
						),
					),
					app.Div().Class("media-content").Body(
						app.Div().Class("content").Body(
							app.Small().Text(a.content[i].ID),
							app.Strong().Text(" "+a.content[i].Title),
							app.Br(),
							app.Text(a.content[i].Type),
						),
						app.Nav().Class("level is-mobile").Body(
							app.Div().Class("level-left").Body(
								app.Div().Class("level-item").Body(
									app.Span().Class("icon is-small").ID(strconv.Itoa(a.content[i].ID)).Body(
										app.I().Class("fas fa-pen"),
									).OnClick(a.onClick),
								),
							),
						),
					),
				),
			)
		}),
	)
}

func (a *article) OnMount(ctx app.Context) {
	go a.doRequest()
}

func (a *article) OnSubmit(ctx app.Context, e app.Event) {
	e.PreventDefault()
	id := app.Window().GetElementByID("id").Get("textContent").String()
	title := app.Window().GetElementByID("title").Get("value").String()
	link := app.Window().GetElementByID("link").Get("value").String()
	image := app.Window().GetElementByID("image").Get("value").String()
	category := app.Window().GetElementByID("category").Get("value").String()
	date := app.Window().GetElementByID("date").Get("value").String()
	creator := app.Window().GetElementByID("creator").Get("value").String()
	contentType := app.Window().GetElementByID("type").Get("value").String()
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
	a.edit = false
	a.doRequest()
}

func (a *article) onClick(ctx app.Context, e app.Event) {
	a.articleID = ctx.JSSrc.Get("id").String()
	if a.edit {
		a.edit = false
		a.loading = false
	} else {
		a.edit = true
		a.loading = true
		go a.doItemRequest(a.articleID)
	}

	a.Update()
}

func (a *article) newArticle(ctx app.Context, e app.Event) {
	a.articleID = ctx.JSSrc.Get("id").String()
	a.edit = true
	a.loading = false
	a.getDefaultItem(a.articleID)

	a.Update()
}

func (a *article) getDefaultItem(articleID string) {
	contentID, _ := strconv.Atoi(articleID)
	a.item = Content{
		ID:      contentID,
		Image:   "https://storage.googleapis.com/hambach/hambach_logo.png",
		Content: "<div class=\"matches\"><p></p></div>",
	}
}

func (a *article) doRequest() {
	resp, err := http.Get("/api/v1/content")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	//Convert the body to type string
	sb := string(body)

	var content []Content
	json.Unmarshal([]byte(sb), &content)

	a.updateResponse(content)
}

func (a *article) updateResponse(content []Content) {
	app.Dispatch(func() {
		a.content = content
		a.nextID = strconv.Itoa(content[0].ID + 1)
		a.Update()
	})
}

func (a *article) doItemRequest(id string) {
	resp, err := http.Get("/api/v1/content?id=" + id)
	if err != nil {
		log.Println(err)
		contentKey := 0
		contentID, _ := strconv.Atoi(id)
		for index, element := range a.content {
			if element.ID == contentID {
				contentKey = index
			}
		}
		a.updateItemResponse(a.content[contentKey])
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	//Convert the body to type string
	sb := string(body)

	var content Content
	json.Unmarshal([]byte(sb), &content)

	a.updateItemResponse(content)
}

func (a *article) updateItemResponse(content Content) {
	app.Dispatch(func() {
		a.item = content
		a.Update()
	})
}
