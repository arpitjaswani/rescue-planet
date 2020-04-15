package controllers

import (
	"admin/models"
	"fmt"
	"net/http"

	"templates"

	"github.com/julienschmidt/httprouter"
)

// Index :
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var data models.TplResponse
	data.Title = "Rescue Planet - Dashboard"
	data.User.UserName = "Arpit J"
	err := templates.Render(w, "dashboard.gohtml", data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
