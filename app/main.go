package main

import (
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &login{})
	app.Route("/home", &home{})
	app.Route("/articles", &articles{})
	app.Route("/metrics", &metrics{})
	app.RouteWithRegexp("^/editor.*", &editor{})

	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite(".", &app.Handler{
		Name:        "Hambach Admin",
		Title:       "Hambach Admin",
		Description: "Admin tools for Spvgg Hambach website",
		Icon: app.Icon{
			Default:    "/web/images/hambach_admin_logo_192.png", // Specify default favicon.
			Large:      "/web/images/hambach_admin_logo_512.png",
			AppleTouch: "/web/images/hambach_admin_logo_192.png", // Specify icon on IOS devices.
		},
		Styles: []string{
			"https://cdn.jsdelivr.net/npm/bulma@0.9.1/css/bulma.min.css",
			"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.2/css/all.min.css",
			"/web/css/main.css",
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
