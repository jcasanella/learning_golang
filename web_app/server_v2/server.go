package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

)

func main() {
	// go get -u github.com/gorilla/mux
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		value, _ := strconv.Atoi(page)
		fmt.Fprintf(w, "You have requested the book: %s on page: %d\n", title, value)
	})

	http.ListenAndServe(":80", r)
}
