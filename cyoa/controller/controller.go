package controller

import (
	"html/template"
	"net/http"

	"github.com/janhaans/golang/cyoa/article"
	"github.com/janhaans/golang/cyoa/templates"
)

func Article(article article.Article) func(http.ResponseWriter, *http.Request) {
	tmpl := template.Must(template.ParseFS(templates.FS, "article.gohtml"))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, article)
	}
}
