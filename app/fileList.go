package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type fileList struct {
	app.Compo

	files []File
}

func (fileList *fileList) OnNav(ctx app.Context) {
	// Launching a new goroutine:
	ctx.Async(func() {
		app_key := app.Getenv("READ_KEY")
		r, err := http.Get("https://api.spvgg-hambach.de/api/v1/files?appkey=" + app_key)
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

		var files []File
		json.Unmarshal([]byte(sb), &files)

		fileList.files = files
		fileList.Update()
	})
}

func (fileList *fileList) Render() app.UI {
	return app.Div().Body(
		app.Button().Class("button is-success").Text("Neuer Upload"),
		app.Br(),
		app.Br(),
		app.Range(fileList.files).Slice(func(i int) app.UI {
			return app.Div().Class("box").Body(
				app.Article().Class("media").Body(
					app.Div().Class("media-left").Body(
						app.Figure().Class("image is-64x64").Body(
							app.Img().Src("https://hambach.s3.eu-central-1.amazonaws.com/"+fileList.files[i].Key),
						),
					),
					app.Div().Class("media-content").Body(
						app.Div().Class("content").Body(
							app.Strong().Text(fileList.files[i].Key),
						),
						app.Nav().Class("level is-mobile").Body(
							app.Div().Class("level-left").Body(
								app.Div().Class("level-item").Body(
									app.Span().Class("icon is-small").ID(fileList.files[i].ID).Body(
										app.I().Class("fas fa-pen"),
									),
								),
							),
						),
					),
				),
			)
		}),
	)
}
