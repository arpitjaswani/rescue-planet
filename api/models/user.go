package models

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

// Users :
var Users = make(map[string]User)

// User :
type User struct {
	UserID   string `json:"userid"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"usertype"`
	Status   string `json:"status"`
}

// AddUser :
func (u User) AddUser() error {
	validUser, msg := validateUserDetails(u)
	if !validUser {
		return errors.New(msg)
	}
	userID, _ := uuid.NewV4()
	u.UserID = userID.String()
	u.Status = "Y"
	// Users[userID.String()] = u

	return nil
}

func validateUserDetails(u User) (valid bool, msg string) {
	if u.UserName == "" {
		msg = "Please enter valid username."
	} else if u.Email == "" {
		msg = "Please enter valid email."
	} else if u.Password == "" {
		msg = "Please enter valid password."
	} else if u.UserType == "" {
		msg = "Please enter valid user type."
	}

	if msg == "" {
		valid = true
	}
	return
}

// GetUsers :
func GetUsers(userID string) ([]User, error) {
	var users []User
	if userID != "" {
		if user, ok := Users[userID]; ok {
			users = append(users, user)
			return users, nil
		}
	}
	for _, user := range Users {
		users = append(users, user)
	}
	return users, nil
}

// UpdateUser :
func (u User) UpdateUser() (User, error) {
	user, ok := Users[u.UserID]
	if !ok {
		return user, errors.New("User not found")
	}
	user.Email = u.Email
	user.Password = u.Password
	user.UserType = u.UserType
	Users[u.UserID] = user
	return user, nil
}

// DeactivateUser :
func DeactivateUser(userID string) (User, error) {
	user, ok := Users[userID]
	if !ok {
		return user, errors.New("User not found")
	}
	user.Status = "N"
	Users[userID] = user
	return user, nil
}
