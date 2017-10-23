// An echo web application.
// problem sheet 2
//17/10/2017
//"fmt"

package main

import (
	
	"net/http"
	
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	http.ServeFile(w, r, "index.html")

}


func guessHandler(w http.ResponseWriter, r *http.Request){
	//http.ServeFile(w, r, "3rdProblem/index.html")
	http.ServeFile(w, r, "guess.html")
}





//main method
func main() {
	http.HandleFunc("/", requestHandler)
	
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)
}