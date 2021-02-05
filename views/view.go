package views

import (
	"html/template"
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
	// calling this function layoutFiles() which gives us a slice
	// but append takes a variadic parameter (...) as a last argument
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

// layoutFiles returns a slice of strings representing
// the layout files used in our application.
func layoutFiles() []string {
	// REMOVED
	// files, err := filepath.Glob("views/layouts/*.gohtml")

	// REPLACED
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}
