package controllers

import (
	"fmt"
	"net/http"

	"templates"

	"github.com/julienschmidt/httprouter"
)

// Index :
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := templates.AdminTpl.ExecuteTemplate(w, "base.gohtml", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
