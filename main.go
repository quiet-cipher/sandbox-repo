package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

const baseDir = "./files"

func fileHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("file")

	base, err := filepath.Abs(baseDir)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	target, err := filepath.Abs(filepath.Join(base, name))
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	if target != base && !strings.HasPrefix(target, base+string(filepath.Separator)) {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	http.ServeFile(w, r, target)
}

func main() {
	http.HandleFunc("/file", fileHandler)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
