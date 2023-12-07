package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/janhaans/golang/cyoa/controllers"
	"github.com/janhaans/golang/cyoa/story"
)

func main() {
	port := flag.Int64("port", 8080, "Port the server is listening at")
	filename := flag.String("file", "gopher.json", "JSON file containing all the chapters of a story")
	flag.Parse()

	r := mux.NewRouter()

	err := story.PublishStory(*filename)
	if err != nil {
		log.Fatalf("Story could not be parsed from JSON file, %v", err)
	}

	r.HandleFunc("/", controllers.IntroHandlerFunc())
	r.HandleFunc("/{chapter}", controllers.ChapterHanderFunc(r))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), r))

}
