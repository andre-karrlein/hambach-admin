package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type fileList struct {
	app.Compo

	files []File
}

func (fileList *fileList) OnNav(ctx app.Context) {
	// Launching a new goroutine:
	ctx.Async(func() {
		sess, _ := session.NewSession(&aws.Config{
			Region: aws.String("eu-central-1")},
		)

		// Create S3 service client
		svc := s3.New(sess)

		resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String("hambach")})
		if err != nil {
			log.Fatalln(err)
		}

		var files []File
		for _, item := range resp.Contents {
			files = append(files, File{
				ID:           *item.ETag,
				Key:          *item.Key,
				LastModified: item.LastModified.String(),
			})
		}

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
