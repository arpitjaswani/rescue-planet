package controllers

import (
	"admin/models"
	"encoding/json"
	"net/http"
	"util"

	"github.com/julienschmidt/httprouter"
)

// AddUser :
func AddUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		util.WebResponse(w, http.StatusForbidden, "Error while fetching request data.")
		return
	}

	err = user.AddUser()
	if err != nil {
		util.WebResponse(w, http.StatusInternalServerError, "Error while adding user.")
		return
	}

	util.WebResponse(w, http.StatusOK, "User added successfully!")
	return
}

// GetUsers :
func GetUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userID := p.ByName("userid")
	users, err := models.GetUsers(userID)
	if len(users) == 0 {
		util.WebResponse(w, http.StatusNotFound, "Users not found.")
		return
	}
	if err != nil {
		util.WebResponse(w, http.StatusNotFound, err.Error())
		return
	}
	util.WebResponse(w, http.StatusOK, users)
	return
}

// UpdateUser :
func UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		util.WebResponse(w, http.StatusForbidden, "Error while fetching request data.")
		return
	}

	if user.UserID == "" {
		util.WebResponse(w, http.StatusForbidden, "Please insert valid user id.")
		return
	}

	_, err = user.UpdateUser()
	if err != nil {
		util.WebResponse(w, http.StatusInternalServerError, "Error while updating user.")
		return
	}

	util.WebResponse(w, http.StatusOK, "User updated successfully!")
	return
}

// DeactivateUser :
func DeactivateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userID := p.ByName("userid")
	_, err := models.DeactivateUser(userID)
	if err != nil {
		util.WebResponse(w, http.StatusNotFound, err.Error())
		return
	}
	util.WebResponse(w, http.StatusOK, "User deactivated successfully!")
	return
}
