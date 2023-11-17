package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/janhaans/golang/nlp"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	// only accept POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	//Do not read more than 1MB from body request
	rdr := io.LimitReader(r.Body, 1_000_000)
	// Read the data from body request
	data, err := io.ReadAll(rdr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	text := string(data)
	//validate
	if len(text) == 0 {
		http.Error(w, "No data in request", http.StatusBadRequest)
		return
	}

	//Get the tokens
	tokens := nlp.Tokenize(text)

	//Set Content-Type in response
	w.Header().Set("Content-Type", "application/json")
	//Create json and write to response
	encoder := json.NewEncoder(w)
	err = encoder.Encode(&tokens)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/tokenize", tokenizeHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("%#v", err)
	}
}
