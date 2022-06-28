package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type articleList struct {
	app.Compo

	content []Content
	nextID  string
}

func (a *articleList) edit(ctx app.Context, e app.Event) {
	articleId := ctx.JSSrc().Get("id").String()
	ctx.Navigate("/editor/" + articleId + "/edit")
}

func (a *articleList) new(ctx app.Context, e app.Event) {
	articleId := ctx.JSSrc().Get("id").String()
	ctx.Navigate("/editor/" + articleId + "/new")
}

func (a *articleList) OnNav(ctx app.Context) {
	// Launching a new goroutine:
	ctx.Async(func() {
		app_key := app.Getenv("READ_KEY")
		r, err := http.Get("https://api.spvgg-hambach.de/api/v1/content?appkey=" + app_key)
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

		var result []Content
		json.Unmarshal([]byte(sb), &result)

		var content []Content
		for _, element := range result {
			content = append(content, element)
		}

		sort.Slice(content, func(i, j int) bool {
			content_i, _ := strconv.Atoi(content[i].ID)
			content_j, _ := strconv.Atoi(content[j].ID)
			return content_i > content_j
		})

		a.content = content
		a.Update()
	})
}

func (a *articleList) Render() app.UI {
	return app.Div().Body(
		app.Button().Class("button is-success").ID(a.nextID).Text("Neuer Artikel").OnClick(a.new),
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
							app.Img().Src(strings.Replace(a.content[i].Image, "https://storage.googleapis.com/hambach/", "https://hambach.s3.eu-central-1.amazonaws.com/", 1)),
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
									app.Span().Class("icon is-small").ID(a.content[i].ID).Body(
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
