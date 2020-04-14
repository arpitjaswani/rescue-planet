package templates

import (
	"html/template"
)

// AdminTpl :
var AdminTpl *template.Template

func init() {
	AdminTpl = template.Must(template.ParseGlob("vendor/templates/admin/*.gohtml"))
}
