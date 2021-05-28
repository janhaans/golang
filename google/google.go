package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	r, err := http.Get("https://google.com")
	if err != nil {
		log.Fatalln(err)
	}

	//Writes the whole response (status, header, body)
	err = r.Write(os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}
}
