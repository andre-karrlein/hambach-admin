package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"google.golang.org/api/iterator"
)

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
	http.HandleFunc("/api/v1/content", contentHandler)

	if err != nil {
		log.Fatal(err)
	}
}

func contentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	contentJSON, err := json.Marshal(loadContent())
	if err != nil {
		log.Fatal(err)
	}

	w.Write(contentJSON)
}

func loadContent() []Content {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	var content []Content

	iter := client.Collection("articles").OrderBy("id", firestore.Desc).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		var c Content
		doc.DataTo(&c)

		content = append(content, c)
	}

	return content
}

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "hambach" // os.Getenv("PROJECT_ID")

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client
}
