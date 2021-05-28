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

	//Low level way to read and output response body
	bs := make([]byte, 99999)
	_, err = r.Body.Read(bs)
	if err != nil && err != io.EOF {
		log.Fatalln(err)
	}

	fmt.Printf("%s\n", string(bs))
}
