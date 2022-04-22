package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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

func (fileList *fileList) OnUpload(ctx app.Context, e app.Event) {
	fileInput := app.Window().GetElementByID("uploadedFile")

	fileInput.Get("files").Call("item", 0).Call("arrayBuffer").Call("then", app.FuncOf(func(v app.Value, x []app.Value) any {
		data := app.Window().Get("Uint8Array").New(x[0])
		dst := make([]byte, data.Get("length").Int())
		app.CopyBytesToGo(dst, data)
		// the data from the file is in dst - do what you want with it
		encoded := base64.StdEncoding.EncodeToString(dst)
		uploadedData := UploadedFile{
			Name: fileInput.Get("name").String(),
			Data: encoded,
		}

		log.Println(uploadedData)

		return nil
	}))
}

func (fileList *fileList) OnDelete(ctx app.Context, e app.Event) {
	id := ctx.JSSrc().Get("id").String()

	ctx.Async(func() {
		app_key := os.Getenv("WRITE_KEY")

		client := &http.Client{}
		// set the HTTP method, url, and request body
		req, err := http.NewRequest(http.MethodDelete, "https://api.spvgg-hambach.de/api/v1/content/"+id+"?appkey="+app_key, bytes.NewBuffer([]byte{}))
		if err != nil {
			panic(err)
		}

		_, err = client.Do(req)
		if err != nil {
			panic(err)
		}
	})
}

func (fileList *fileList) OnChoose(ctx app.Context, e app.Event) {
	var file File
	for _, value := range fileList.files {
		if value.ID == ctx.JSSrc().Get("id").String() {
			file = value
		}
	}
	ctx.SetState("file", file)
}

func (fileList *fileList) Render() app.UI {
	return app.Div().Body(
		app.Input().Type("file").ID("uploadedFile").Name("uploadedFile"),
		app.Button().Class("button is-success").Text("Neuer Upload").OnClick(fileList.OnUpload),
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
						app.Div().Body(
							app.Span().Class("is-small").Body(
								app.Text("LINK: https://hambach.s3.eu-central-1.amazonaws.com/"+fileList.files[i].Key),
							),
						),
						app.Nav().Class("level is-mobile").Body(
							app.Div().Class("level-left").Body(
								app.Div().Class("level-item").Body(
									app.Span().Class("icon is-small").ID(fileList.files[i].Key).Body(
										app.I().Class("fas fa-trash"),
									).OnClick(fileList.OnDelete),
								),
							),
						),
					),
				),
			)
		}),
	)
}
