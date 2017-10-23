// An echo web application.
// problem sheet 2
//17/10/2017
//"fmt"

package main

import (
	
	"net/http"
	"html/template"
	
)

type message struct {
	MyMessage string
	
}



func requestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	http.ServeFile(w, r, "index.html")

	

}


func guessHandler(w http.ResponseWriter, r *http.Request){
	msm :="Guess a number between 1 and 20"
	

	//http.ServeFile(w, r, "3rdProblem/index.html")
	http.ServeFile(w, r, "guess.html")

	t, _ := template.ParseFiles("guess.tmpl")

	t.Execute(w, &message{MyMessage:msm})
}





//main method
func main() {
	http.HandleFunc("/", requestHandler)
	
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)
}