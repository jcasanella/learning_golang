package main

import (
	"fmt"
	"log"
	"net/http"

)

func main() {

	// Create handler which receives all incoming HTTP connections from browsers.
	// Response writer used to print the html
	// http.Request contains all the info about this Http request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello you've requested from: %s</h1>\n", r.URL.Path)
		if name := r.URL.Query().Get("name"); name != "" {
			fmt.Fprintf(w, "What's up %s", name)
		}
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<h1>This is a test page</h1>")
		if name := r.URL.Query().Get("name"); name != "" {
			fmt.Fprintf(w, "Name is: <strong>%s</strong>", name)
		}
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// To accept connections, a server needs to listen on a port to pass connections to request handler
	log.Fatal(http.ListenAndServe(":80", nil))
}
