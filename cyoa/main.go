package main

import (
	"log"
	"net/http"

	"github.com/janhaans/golang/cyoa/article"
	"github.com/janhaans/golang/cyoa/controller"
)

func main() {
	mux := http.NewServeMux()

	articles, err := article.ReadArticles("gopher.json")
	if err != nil {
		log.Fatalf("%#v", err)
	}
	for path, article := range articles {
		if path == "intro" {
			mux.HandleFunc("/", controller.Article(article))
		} else {
			mux.HandleFunc("/"+path, controller.Article(article))
		}
	}

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("%#v", err)
	}

}
