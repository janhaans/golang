package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/janhaans/golang/nlp"
)

type Server struct {
	logger *log.Logger
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func (s *Server) tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	// only accept POST requests
	//if r.Method != http.MethodPost {
	//	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	//	return
	//}
	//defer r.Body.Close()

	//Do not read more than 1MB from body request
	rdr := io.LimitReader(r.Body, 1_000_000)
	// Read the data from body request
	data, err := io.ReadAll(rdr)
	if err != nil {
		s.logger.Printf("error - cannot read %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	text := string(data)
	//validate
	if len(text) == 0 {
		s.logger.Printf("error - no data %s", err)
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
	logger := log.New(log.Writer(), "nlp", log.Ldate|log.Ltime|log.Llongfile)
	s := Server{logger}
	r := mux.NewRouter()
	r.HandleFunc("/health", s.healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/tokenize", s.tokenizeHandler).Methods(http.MethodPost)
	http.Handle("/", r)
	addr := ":8080"
	s.logger.Printf("Server starting at port %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("%#v", err)
	}
}
