// An echo web application.
// problem sheet 2
//17/10/2017
//"fmt"
//"math/rand"

/*
"time"
	"strconv"
	"html/template"
	"math/rand"
	*/

package main

import (
	
	"net/http"
	"html/template"
	//rand.Intn(20-1) -- to generate a number between 1 and 20
	
)

type message struct {
	MyMessage string
	
}


func requestHandler(w http.ResponseWriter, r *http.Request) {	
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")

}


func guessHandler(w http.ResponseWriter, r *http.Request){

	msm :=&message{MyMessage: "Guess a number between 1 and 20"}

	 t, _ := template.ParseFiles("guess.tmpl") 

	t.Execute(w, msm)
	
}//end guesshandler


//main method
func main() {
	http.HandleFunc("/", requestHandler)
	
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)
}