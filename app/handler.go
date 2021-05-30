package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

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
