package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var (
	fsDir       = http.FileServer(http.Dir(filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "enahs", "gotemplatesbyexample", "assets")))
	tmplatesDir = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "enahs", "gotemplatesbyexample", "templates")
)

func main() {

	http.Handle("/assets/", http.StripPrefix("/assets/", fsDir))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index := filepath.Join(tmplatesDir, "/index.html.tmpl")
		t, err := template.ParseFiles(index)
		if err != nil {
			w.Write([]byte("could not parse template!" + err.Error()))
		}
		t.Execute(w, nil)
	})
	http.ListenAndServe(":8080", nil)
}
