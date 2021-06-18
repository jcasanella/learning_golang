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
		fmt.Fprintf(w, "Hello you've requested from: %s\n", r.URL.Path)
	})

	// To accept connections, a server needs to listen on a port to pass connections to request handler
	log.Fatal(http.ListenAndServe(":80", nil))
}
