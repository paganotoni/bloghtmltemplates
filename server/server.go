package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"text/template"
	"time"
)

// The layout template.
var layout *template.Template

func renderPost(w http.ResponseWriter, r *http.Request) {
	// Some caching layer here.
	if layout == nil {
		dd, err := ioutil.ReadFile("./templates/layout.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		layout, err = template.New("layout").Parse(string(dd))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	dd, err := ioutil.ReadFile(filepath.Join("./templates/posts/", r.URL.Path))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tt := fmt.Sprintf(`{{define "content"}}%s{{end}}`, dd)
	tmpl, err := template.Must(layout.Clone()).Parse(tt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Year  string
		Title string
	}{
		Year:  time.Now().Format("2006"),
		Title: "The Blog",
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", renderPost)

	return mux
}
