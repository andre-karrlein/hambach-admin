package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type navbar struct {
	app.Compo

	active string
}

// Status reflects the logged in status
type Status struct {
	LoggedIn bool `json:"loggedIn"`
}

type home struct {
	app.Compo
}

type article struct {
	app.Compo

	content   []Content
	edit      bool
	articleID string
	item      Content
}

const hash = "$2y$14$7aNuDEs7G6KxyYZLShEHlOpY4cjxV4kizm3noGFNBW11dvJdgtp3G"

// GetUsers function to get all available users with password
func GetUsers() map[string]string {
	return map[string]string{
		"akarrlein": hash,
		"pgeissler": hash,
	}
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

func (h *home) Render() app.UI {
	var status Status
	app.SessionStorage.Get("status", &status)

	if status.LoggedIn != true {
		app.Navigate("/login")
	}

	return app.Div().Body(
		&navbar{},
		app.Section().Class("section").Body(
			app.Body().Class("has-navbar-fixed-top").Body(
				app.Div().Class("container").Body(
					&article{},
				),
			),
		),
	)
}

func (n *navbar) OnClick(ctx app.Context, e app.Event) {
	status := Status{LoggedIn: false}
	app.SessionStorage.Set("status", status)
	app.Navigate("/login")
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

func (a *article) Render() app.UI {
	if a.edit {
		go a.doItemRequest(a.articleID)

		return app.Div().Body(
			app.Button().Class("button").Text("Zurück").OnClick(a.onClick),
			app.Br(),
			app.Br(),
			app.Form().Class("box").OnSubmit(a.OnSubmit).Body(
				app.Small().Text(a.item.ID).ID("id"),
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
					app.Label().Class("label").Text("Ersteller ID"),
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
	go a.doRequest("")

	return app.Div().Body(
		app.Range(a.content).Slice(func(i int) app.UI {

			return app.Div().Class("box").Body(
				app.Article().Class("media").Body(
					app.Div().Class("media-left").Body(
						app.Figure().Class("image is-64x64").Body(
							app.Img().Src(a.content[i].Image),
						),
					),
					app.Div().Class("media-content").Body(
						app.Div().Class("content").Body(
							app.Strong().Text(a.content[i].Title),
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

func (a *article) OnSubmit(ctx app.Context, e app.Event) {
	e.PreventDefault()
	/*id := app.Window().GetElementByID("id").Get("textContent").Int()
	title := app.Window().GetElementByID("title").Get("value").String()
	link := app.Window().GetElementByID("link").Get("value").String()
	image := app.Window().GetElementByID("image").Get("value").String()
	category := app.Window().GetElementByID("category").Get("value").String()
	date := app.Window().GetElementByID("date").Get("value").String()
	creator := app.Window().GetElementByID("creator").Get("value").String()
	contentType := app.Window().GetElementByID("type").Get("value").String()
	content := app.Window().GetElementByID("content").Get("textContent").String()
	active := app.Window().GetElementByID("active").Get("value").String()

	article := Content{
		id,
		title,
		date,
		category,
		contentType,
		image,
		creator,
		content,
		active,
		link,
	}

	*/
}

func (a *article) onClick(ctx app.Context, e app.Event) {
	a.articleID = ctx.JSSrc.Get("id").String()
	if a.edit {
		a.edit = false
	} else {
		a.edit = true
	}

	a.Update()
}

func (a *article) doRequest(uri string) {
	resp, err := http.Get("/api/v1/content" + uri)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)

	var content []Content
	json.Unmarshal([]byte(sb), &content)

	a.updateResponse(content)
}

func (a *article) updateResponse(content []Content) {
	app.Dispatch(func() { // Ensures response field is updated on UI goroutine.
		a.content = content
		a.Update()
	})
}

func (a *article) doItemRequest(id string) {
	resp, err := http.Get("/api/v1/content?id=" + id)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)

	var content Content
	json.Unmarshal([]byte(sb), &content)

	a.updateItemResponse(content)
}

func (a *article) updateItemResponse(content Content) {
	app.Dispatch(func() { // Ensures response field is updated on UI goroutine.
		a.item = content
		a.Update()
	})
}

func main() {
	app.Route("/", &home{})
	app.Route("/login", &login{})
	app.Run()
}
