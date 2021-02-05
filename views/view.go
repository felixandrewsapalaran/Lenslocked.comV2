package views

import "html/template"

// NewView parse all the files necessary and return them to us
func NewView(layout string, files ...string) *View {

	files = append(files,
		"views/layouts/footer.gohtml",
		"views/layouts/navbar.gohtml", // add navbar template
		"views/layouts/bootstrap.gohtml",
	)

	// which ever `yield` pass in is the one to get use
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	// return pointer to a view
	return &View{
		Template: t,
		Layout:   layout,
	}
}

// View contains our logic template
type View struct {
	Template *template.Template
	// this Layout is going to store the layout we want to render
	Layout string
}
