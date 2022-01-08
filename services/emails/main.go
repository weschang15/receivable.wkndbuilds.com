// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/weschang15/receivable.wkndbuilds.com/services/emails/handlers"
)

func main() {
	l := log.New(os.Stdout, "receivable-wnkdbuilds-emails", log.LstdFlags)
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.EmailHandler)

	// retrieve environment variable port value
	p := os.Getenv("PORT")

	if p == "" {
		p = "8080"
		l.Printf("defaulting to port %s", p)
	}

	// create a new http Server to override default properties
	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + p,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	l.Printf("listening on port %s", p)
	if err := srv.ListenAndServe(); err != nil {
		l.Fatalf("Error starting server: %s\n", err)
	}
}
