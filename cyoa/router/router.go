package router

import (
	"github.com/gorilla/mux"
	"github.com/janhaans/golang/cyoa/controllers"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.IntroHandlerFunc())
	r.HandleFunc("/{chapter}", controllers.ChapterHanderFunc(r))

	return r
}
