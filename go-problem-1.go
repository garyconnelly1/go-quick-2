// An echo web application.
// problem sheet 2

// "fmt"

package main

import (
	
	
	
	"net/http"
	
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	//fmt.Fprintln(w, "Guessing Game")
	http.ServeFile(w, r, "3rdProblem/index.html")

}

//main method
func main() {
	http.HandleFunc("/guess", requestHandler)
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}