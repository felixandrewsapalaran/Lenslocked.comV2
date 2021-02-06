package controllers

import (
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

// New use to utilize the `NewView` view
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	// check for error in case if not rendering correctly
	err := u.NewView.Render(w, nil)
	if err != nil {
		panic(err)
	}

}
