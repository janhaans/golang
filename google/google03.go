package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	r, err := http.Get("https://google.com")
	if err != nil {
		log.Fatalln(err)
	}
	io.Copy(os.Stdout, r.Body)
}
