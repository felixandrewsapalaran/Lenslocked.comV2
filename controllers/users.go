package controllers

import (
	"fmt"
	"net/http"

	"Lenslocked.com/views"
)

//NewUsers is used to create a new Users controllers.
// This function will panic if the templates are not
// parsed correctly, and should only be use during
// initial setup
func NewUsers() *Users {
	return &Users{
		// setting up our template file
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

// Users contains a pointer to our type View
type Users struct {
	// adding a view type
	NewView *views.View
}

// New use to render the form where a user can
// create a new user account
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	// check for error in case if not rendering correctly
	err := u.NewView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

// Create is used to process the signup form when a user
// submits it. This is used to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a temporary response. ")
}
