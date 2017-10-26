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
	"math/rand"
	"time"
	
	"strconv"
	
	
	
)
//rand.Intn(20-1) -- to generate a number between 1 and 20

//message struvt
type message struct {
	MyMessage string
	
}

//form data struct




func requestHandler(w http.ResponseWriter, r *http.Request) {	
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")

}


func guessHandler(w http.ResponseWriter, r *http.Request){

	msm :=&message{MyMessage: "Guess a number between 1 and 20"}

	

//parse the form data


	
	

//generate random number




	target:=rand.Intn(20-1)

	var ck, err = r.Cookie("target")//gets cookie called count

	

	if err == nil{
		//if we could read it ,try to convert its value to an int
		target, _ = strconv.Atoi(ck.Value)
		
		if target ==0{
			target = rand.Intn(20-1)
		}//end if
		
	}//end if

	expires := time.Now().AddDate(1, 0, 0) // cookie will expire after 1 year

    ck = &http.Cookie{
		Name: "target",
        Value: strconv.Itoa(target),
        Path: "/",
        Expires: expires,
    }

	//set the cookie
	http.SetCookie(w,ck)
	



	 t, _ := template.ParseFiles("guess.tmpl") 

	t.Execute(w, msm)

	

	
}//end guesshandler


//main method
func main() {
	http.HandleFunc("/", requestHandler)
	
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)
}

//the following is the code for the data parsing.. kept getting a runtime error with is
//it would all go in the guessHandler function

/*
type formData struct {
	guess1 string
}

//parse the form data
	r.ParseForm()
	//guess1,_ := strconv.Atoi(r.FormValue("Guess"))
    user := &formData{guess1: r.Form["guess"][0]}

	t2, _ := template.ParseFiles("guess.tmpl") 

	t2.Execute(w, user)
*/
