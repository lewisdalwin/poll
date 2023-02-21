// Filename: cmd/web/routes.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func(app *application) routes() *httprouter.Router {
	router := httprouter.New()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
  router.Handler(http.MethodGet,"/static/*filepath", http.StripPrefix("/static", fileServer))
	router.HandlerFunc(http.MethodGet,"/", app.home)
	router.HandlerFunc(http.MethodGet, "/about", app.about)
	router.HandlerFunc(http.MethodGet,"/question/create", app.questionCreate)
	router.HandlerFunc(http.MethodPost, "/question/create", app.questionCreatePost)
	// return the router
	return router
}