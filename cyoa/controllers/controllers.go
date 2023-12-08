package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/janhaans/golang/cyoa/story"
	"github.com/janhaans/golang/cyoa/templates"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFS(templates.FS, "chapter.gohtml"))
}

func IntroHandlerFunc() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, story.PublishedStory["intro"])
		if err != nil {
			log.Printf("Template execution failed %v\n", err)
			http.Error(w, "Template execution failed", http.StatusBadGateway)
		} else {
			log.Println("Chapter intro served")
		}
	}
}

func ChapterHanderFunc(r *mux.Router) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		path := vars["chapter"]
		if chapter, ok := story.PublishedStory[path]; ok {
			err := tmpl.Execute(w, chapter)
			if err != nil {
				log.Printf("Template execution failed %v\n", err)
				http.Error(w, "Template execution failed", http.StatusBadGateway)
			} else {
				log.Printf("Chapter %s served\n", path)
			}
			return
		}
		log.Printf("There is not chapter %s\n", path)
		http.NotFound(w, r)
	}
}
