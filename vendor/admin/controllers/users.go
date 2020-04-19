package controllers

import (
	"admin/models"
	"data"
	"encoding/json"
	"fmt"
	"net/http"
	"templates"
	"util"

	"golang.org/x/crypto/bcrypt"

	"github.com/julienschmidt/httprouter"
)

// Register :
func Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	err := templates.Render(w, "registration.gohtml", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

// Adopt :
func Adopt(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	err := templates.Render(w, "dashboard.gohtml", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

// AddUser :
func AddUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var data models.TplResponse
	data.Title = "Rescue Planet - Dashboard"
	err := templates.Render(w, "dashboard.gohtml", data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

// GetUsers :
func GetUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	/*userID := p.ByName("userid")
	users, err := models.GetUsers(userID)
	if len(users) == 0 {
		util.WebResponse(w, http.StatusNotFound, "Users not found.")
		return
	}
	if err != nil {
		util.WebResponse(w, http.StatusNotFound, err.Error())
		return
	}*/
	util.WebResponse(w, http.StatusOK, "users")
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

// Login :
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := templates.Render(w, "login.gohtml", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func CheckLogin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var userDetails data.UserClaims
	err := json.NewDecoder(r.Body).Decode(&userDetails)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		fmt.Println(err.Error())
		util.WebResponse(w, http.StatusBadRequest, "invaild input json")
		return
	}

	if userDetails.Email == "" || userDetails.Password == "" {
		util.WebResponse(w, http.StatusBadRequest, "invaild request")
		return
	}

	user, err := data.GetUser(userDetails.Email, "admin")
	if user.Email == "" {
		util.WebResponse(w, http.StatusNotFound, "Users not found.")
		return
	}
	if err != nil {
		util.WebResponse(w, http.StatusNotFound, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDetails.Password))
	if err != nil {
		util.WebResponse(w, http.StatusUnauthorized, "Unauthorized access")
		return
	}
	userDetails.Password = ""
	// generating the JWT token for the user
	token, err := data.GenerateToken(userDetails)
	if err != nil {
		fmt.Println(err.Error())
		util.WebResponse(w, http.StatusBadGateway, "Unable to generate token")
		return
	}

	// Set token on cookies
	data.SetTokenCookies(w, token)

	var data models.TplResponse
	data.Title = "Rescue Planet - Dashboard"
	data.User.UserName = userDetails.Username
	err = templates.Render(w, "dashboard.gohtml", data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
