package server

import (
	"blog/content"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
	"time"
)

// The layout template.
var (
	layout     *template.Template
	commonData = struct {
		Year  string
		Title string
	}{
		Year:  time.Now().Format("2006"),
		Title: "The Blog",
	}
)

func renderPost(w http.ResponseWriter, r *http.Request) {
	// Some caching layer here.
	dd, err := content.All.ReadFile("layout.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	layout, err = template.New("layout").Parse(string(dd))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dd, err = content.All.ReadFile(filepath.Base(r.URL.Path))
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

	if err := tmpl.Execute(w, commonData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
