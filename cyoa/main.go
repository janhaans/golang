package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/janhaans/golang/cyoa/controller"
	"github.com/janhaans/golang/cyoa/story"
)

func main() {
	filename := flag.String("file", "gopher.json", "JSON file containing all the chapters of a story")
	flag.Parse()
	mux := http.NewServeMux()

	story, err := story.GetStory(*filename)
	if err != nil {
		log.Fatalf("%#v", err)
	}

	for path, chapter := range story {
		if path == "intro" {
			mux.HandleFunc("/", controller.Chapter(chapter))
		} else {
			mux.HandleFunc("/"+path, controller.Chapter(chapter))
		}
	}

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("%#v", err)
	}

}
