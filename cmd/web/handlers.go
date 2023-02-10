package main

import (
	"net/http"
)

// home displays the home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	RenderTemplate(w, "home.page.tmpl")
}

// about displays the about us page
func (app *application) about(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.page.tmpl")

}

// createFeedback creates a new feedback
func (app *application) create(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "poll.page.tmpl")
}
