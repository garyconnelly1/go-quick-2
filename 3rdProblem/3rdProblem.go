// An echo web application.
// problem sheet 2
//17/10/2017

package main

import (
	"fmt"
	"net/http"
	
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintln(w, r, "index.html")

}

//main method
func main() {
	http.HandleFunc("/", requestHandler)
	
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)
}