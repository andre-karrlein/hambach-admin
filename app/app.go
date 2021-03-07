package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

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

	content []Content
	edit    bool
	article string
}

const hash = "$2y$14$7aNuDEs7G6KxyYZLShEHlOpY4cjxV4kizm3noGFNBW11dvJdgtp3G"

// GetUsers function to get all available users with password
func GetUsers() map[string]string {
	return map[string]string{
		"akarrlein": hash,
		"pgeissler": hash,
	}
}

type experience struct {
	Role      string `json:"role"`
	Company   string `json:"company"`
	City      string `json:"city"`
	Timeframe string `json:"timeframe"`
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
			app.Body().Body(
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
	return app.Nav().Class("navbar is-success").Body(
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
	if !a.edit {
		go a.doRequest()

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
										app.Span().Class("icon is-small").Body(
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
	return app.Div().Text("Edit")
}

func (a *article) onClick(ctx app.Context, e app.Event) {
	a.edit = true

	a.Update()
}

func (a *article) doRequest() {
	resp, err := http.Get("/api/v1/content")
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

func main() {
	app.Route("/", &home{})
	app.Route("/login", &login{})
	app.Run()
}
