package main

type File struct {
	ID           string `json:"id"`
	Key          string `json:"key"`
	LastModified string `json:"last_modified"`
}

type UploadedFile struct {
	Name string `json:"name"`
	Data string `json:"data"`
}
