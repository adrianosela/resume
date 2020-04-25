package main

import (
	"log"
	"net/http"
	"os"
)

const (
	resume      = "resume.pdf"
	defaultPort = ":80"
)

func servePDF(pdf string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, pdf)
	})
}

func main() {
	// When running on Google App Engine, the PORT
	// env variable is set by the runtime
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	err := http.ListenAndServe(port, servePDF(resume))
	if err != nil {
		log.Fatal(err)
	}
}
