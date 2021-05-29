package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type articleList struct {
	app.Compo

	content []Content
	nextID  string
}

func (a *articleList) edit(ctx app.Context, e app.Event) {
	articleId := ctx.JSSrc.Get("id").String()
	ctx.Navigate("/editor/" + articleId)
}

func (a *articleList) OnNav(ctx app.Context) {
	// Launching a new goroutine:
	ctx.Async(func() {
		r, err := http.Get("/api/v1/content") // #TODO
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

			var content []Content
			json.Unmarshal([]byte(sb), &content)

			a.content = content
			a.nextID = strconv.Itoa(content[0].ID + 1)
			a.Update()
		})
	})
}

func (a *articleList) Render() app.UI {
	return app.Div().Body(
		app.Button().Class("button is-success").ID(a.nextID).Text("Neuer Artikel").OnClick(a.edit),
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
									).OnClick(a.edit),
								),
							),
						),
					),
				),
			)
		}),
	)
}
