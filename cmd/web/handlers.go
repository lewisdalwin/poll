package main

import (
	"fmt"
	"net/http"
)

// home displays the home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.page.tmpl")
}

// about displays the about us page
func (app *application) about(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.page.tmpl")

}

// display the question form
func (app *application) questionCreate(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "poll.page.tmpl")
}

// submit the question form
func (app *application) questionCreatePost(w http.ResponseWriter, r *http.Request) {
	// test the insert functionality
	question := "How are you doing Dalwin?"
	id, err := app.questions.Insert(question)
	if err != nil {
		// server error then return
		return
	}
	// get the id
	fmt.Fprintf(w, "The question has id: %d\n", id)
}