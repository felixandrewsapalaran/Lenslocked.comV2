package views

/*This file will handle our view logic*/

import "html/template"

// NewView parse all the files necessary and return them to us
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")

	// Parse the files
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	// return pointer to a view
	return &View{
		Template: t,
	}
}

// View contains our logic template
type View struct {
	Template *template.Template
}
