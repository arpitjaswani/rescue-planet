package main

import (
	"data"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"util"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/julienschmidt/httprouter"
)

// GenerateToken :
func GenerateToken(pathType string, h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var userDetails data.UserClaims
		err := json.NewDecoder(r.Body).Decode(&userDetails)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			fmt.Println(err.Error())
			util.WebResponse(w, http.StatusBadRequest, "invaild input json")
			return
		}

		if userDetails.Email == "" || userDetails.Username == "" || userDetails.Password == "" {
			util.WebResponse(w, http.StatusBadRequest, "invaild request")
			return
		}

		dbUser, _ := data.GetUser(userDetails.Email, pathType)
		if err != nil {
			fmt.Println(err.Error())
		}

		if dbUser.Email != "" {
			util.WebResponse(w, http.StatusBadRequest, "Email already register")
			return
		}

		// Generate Hash password
		hashPass, err := bcrypt.GenerateFromPassword([]byte(userDetails.Password), bcrypt.MinCost)
		if err != nil {
			util.WebResponse(w, http.StatusBadGateway, "Error While Hashing Password, Try Again")
			return
		}
		userDetails.Password = ""
		userDetails.UserType = pathType

		// generating the JWT token for the user
		token, err := data.GenerateToken(userDetails)
		if err != nil {
			fmt.Println(err.Error())
			util.WebResponse(w, http.StatusBadGateway, "Unable to generate token")
			return
		}

		// Set token on cookies
		data.SetTokenCookies(w, token)

		r.Header.Set("Authorization", "bearer "+token)
		data.UserRegister(userDetails, hashPass)
		// Proceed
		h(w, r, ps)
		return
	}
}

func IsAuthorized(pathType string, h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var userClaims data.StandardClaims
		if r.Header["Authorization"] != nil {
			tokenStrings := strings.Split(r.Header["Authorization"][0], " ")
			fmt.Println("token: ", tokenStrings[1])
			if strings.ToLower(tokenStrings[0]) != "bearer" || tokenStrings[1] == "" {
				util.WebResponse(w, http.StatusBadRequest, "Invaild Access")
				return
			}

			tkn, err := jwt.ParseWithClaims(tokenStrings[1], &userClaims, func(token *jwt.Token) (interface{}, error) {
				return data.MySigningKey, nil
			})

			if err != nil || !tkn.Valid {
				fmt.Println(err.Error())
				util.WebResponse(w, http.StatusUnauthorized, "Unauthorized access")
				return
			}
			fmt.Println(userClaims)
		} else {
			util.WebResponse(w, http.StatusUnauthorized, "Not Aauthorized")
			return
		}

		// Proceed
		h(w, r, ps)
		return
	}
}
