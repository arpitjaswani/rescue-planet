package templates

import (
	"html/template"
	"io"
)

// adminTpl :
var adminTpl *template.Template

func init() {
	adminTpl = template.Must(template.ParseGlob("vendor/templates/admin/*.gohtml"))
}

// Render :
func Render(wr io.Writer, name string, data interface{}) error {
	err := adminTpl.ExecuteTemplate(wr, name, data)
	if err != nil {
		// sent another resposne here
	}
	return err
}
