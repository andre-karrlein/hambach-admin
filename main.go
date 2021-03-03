package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// The main function is the entry of the server. It is where the HTTP handler
// that serves the UI is defined and where the server is started.
//
// Note that because main.go and app.go are built for different architectures,
// this main() function is not in conflict with the one in
// app.go.
func main() {
	// app.Handler is a standard HTTP handler that serves the UI and its
	// resources to make it work in a web browser.
	//
	// It implements the http.Handler interface so it can seamlessly be used
	// with the Go HTTP standard library.
	http.Handle("/", &app.Handler{
		Name:        "Hambach Admin",
		Title:       "Hambach Admin",
		Description: "Admin tools for Spvgg Hambach website",
		Icon: app.Icon{
			Default:    "https://storage.googleapis.com/hambach/IMG_0265.JPG", // Specify default favicon.
			AppleTouch: "https://storage.googleapis.com/hambach/IMG_0265.JPG", // Specify icon on IOS devices.
		},
		Styles: []string{
			"https://cdn.jsdelivr.net/npm/bulma@0.9.1/css/bulma.min.css",
			"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.2/css/all.min.css",
			"/web/css/main.css",
		},
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
