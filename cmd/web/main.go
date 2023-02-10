package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

// Share data across our handlers
type application struct {}

func main() {
	// configure our server
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// share data across our handlers
	app := &application{}
	// create and start a custom web server
	log.Printf("starting server on %s", *addr)
	srv := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err := srv.ListenAndServe()
	log.Fatal(err)
}
