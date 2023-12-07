package controllers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/janhaans/golang/cyoa/story"
	"github.com/janhaans/golang/cyoa/templates"
)

func IntroHandlerFunc() func(http.ResponseWriter, *http.Request) {
	tmpl := template.Must(template.ParseFS(templates.FS, "chapter.gohtml"))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, story.PublishedStory["intro"])
	}
}

func ChapterHanderFunc(r *mux.Router) func(http.ResponseWriter, *http.Request) {
	tmpl := template.Must(template.ParseFS(templates.FS, "chapter.gohtml"))
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		chapter := vars["chapter"]
		if _, ok := story.PublishedStory[chapter]; !ok {
			http.NotFound(w, r)
		} else {
			tmpl.Execute(w, story.PublishedStory[chapter])
		}
	}
}
