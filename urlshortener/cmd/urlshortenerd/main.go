package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/janhaans/golang/urlshortener"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	mux = urlshortener.MapHandler(pathsToUrls, mux)
	mux, err := urlshortener.YAMLHandler([]byte(yaml), mux)
	if err != nil {
		log.Fatalf("%#v", err)
	}
	fmt.Println("Starting the server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("%#v", err)
	}

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello World!</h1>")
}
