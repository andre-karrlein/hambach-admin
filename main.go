package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

type article struct {
	ID       int    `firestore:"id,omitempty"`
	Title    string `firestore:"title,omitempty"`
	Date     string `firestore:"date,omitempty"`
	Category string `firestore:"category,omitempty"`
	Type     string `firestore:"type,omitempty"`
	Image    string `firestore:"image,omitempty"`
	Creator  string `firestore:"creator,omitempty"`
	Content  string `firestore:"content,omitempty"`
	Active   string `firestore:"active,omitempty"`
	Link     string `firestore:"link,omitempty"`
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
	
	http.HandleFunc("/api/v1/content", contentHandler)
	http.HandleFunc("/api/v1/content/save", saveHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		id := r.FormValue("id")
		title := r.FormValue("title")
		date := r.FormValue("date")
		category := r.FormValue("category")
		contentType := r.FormValue("contentType")
		image := r.FormValue("image")
		creator := r.FormValue("creator")
		content := r.FormValue("content")
		active := r.FormValue("active")
		link := r.FormValue("link")

		articleID, _ := strconv.Atoi(id)

		article := article{
			ID:       articleID,
			Title:    title,
			Date:     date,
			Category: category,
			Image:    image,
			Type:     contentType,
			Creator:  creator,
			Content:  content,
			Active:   active,
			Link:     link,
		}

		saveArticle(article)

		w.WriteHeader(http.StatusCreated)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func contentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		contentJSON, err := json.Marshal(loadContent())
		if err != nil {
			log.Fatal(err)
		}

		w.Write(contentJSON)
		return
	}

	id := keys[0]

	contentJSON, err := json.Marshal(loadArticle(id))
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

func loadArticle(id string) Content {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	result, err := client.Collection("articles").Doc(id).Get(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	var c Content
	result.DataTo(&c)

	return c
}

func saveArticle(article article) {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	_, err := client.Collection("articles").Doc(strconv.Itoa(article.ID)).Set(ctx, article)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
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
