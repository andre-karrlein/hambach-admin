package main

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
