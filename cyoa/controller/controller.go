package controller

import (
	"html/template"
	"net/http"

	"github.com/janhaans/golang/cyoa/story"
	"github.com/janhaans/golang/cyoa/templates"
)

func Chapter(chapter story.Chapter) func(http.ResponseWriter, *http.Request) {
	tmpl := template.Must(template.ParseFS(templates.FS, "article.gohtml"))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, chapter)
	}
}
