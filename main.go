package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template

func main() {
	// declare err here
	var err error
	// home.gohtml is the template we are using to render our page
	// this way it assumes both variables homeTemplate and err are already declared
	// using this approach prevents scoping problem especially for "homeTemplate"
	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	// check if our template load properly
	if err != nil {
		// if error occurs for loading our template panic the error
		// the program will stops here if there's a problem
		panic(err)
	}
	// This server wont start if there's problem loading the template
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}

// home handler
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// updated our function to use our template
	// by executing the homeTemplate
	if err := homeTemplate.Execute(w, nil); err != nil { // wrapping it using "if err equal type" thing
		// if there an error executing the template then panic the error
		panic(err)
	}
}

// contact handler
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email to "+
		"<a href=\"mailto:support@usegolang.com\">support@usegolang.com</a>.")
}
