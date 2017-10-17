// An echo web application.
// problem sheet 2

package main

import (
	"fmt"
	"net/http"
	
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintln(w, "<h1>Guessing Game<h1>")

}

//main method
func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}