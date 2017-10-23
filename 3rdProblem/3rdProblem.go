// An echo web application.
// problem sheet 2
//17/10/2017
//"fmt"
//"math/rand"

package main

import (
	
	"net/http"
	"html/template"
	"math/rand"
	"time"
	"strconv"
	
	
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
	msm :="Guess a number between 1 and 20"

	target := rand.Intn(20-1)


//adapted from https://stackoverflow.com/questions/12130582/setting-cookies-in-golang-net-http
	 expires := time.Now().AddDate(1, 0, 0) // cookie will expire after 1 year

    ck := &http.Cookie{
        Name: "target",
        Value: strconv.Itoa(target),
        //Path: "/",
        Expires: expires,
    }

	//set the cookie
	http.SetCookie(w,ck)
	

	//http.ServeFile(w, r, "3rdProblem/index.html")
	http.ServeFile(w, r, "guess.html")

	t, _ := template.ParseFiles("guess.tmpl")

	t.Execute(w, &message{MyMessage:msm})
}//end guesshandler





//main method
func main() {
	http.HandleFunc("/", requestHandler)
	
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)
}