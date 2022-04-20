package main

type File struct {
	ID           string
	Key          string
	LastModified string
}

type UploadedFile struct {
	Name string `json:"name"`
	Data string `json:"data"`
}
