package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	r, err := http.Get("https://google.com")
	if err != nil {
		log.Fatalln(err)
	}

	//High level way to read and output response body
	bs, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", string(bs))

}
