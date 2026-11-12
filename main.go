package main

import (
	"log"
	"net/http"
	"path/filepath"
)

const baseDir = "./files"

func fileHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("file")
	path := filepath.Join(baseDir, name)
	http.ServeFile(w, r, path)
}

func main() {
	http.HandleFunc("/file", fileHandler)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
