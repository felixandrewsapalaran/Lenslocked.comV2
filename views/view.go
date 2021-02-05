package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	// LayoutDir is the directory where we store our layouts in
	LayoutDir string = "views/layouts/"
	// TemplateExt is the extension of our files
	TemplateExt string = ".gohtml"
)

// NewView parse all the files necessary and return them to us
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

// View contains our logic template
type View struct {
	Template *template.Template
	Layout   string
}

// Render is used to render the view with predefined layout.
func (v *View) Render(w http.ResponseWriter, data interface{}) error { // data the is the data we want to pass in when we execute our template
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// layoutFiles returns a slice of strings representing
// the layout files used in our application.
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}
