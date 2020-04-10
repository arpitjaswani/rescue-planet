package models

import (
	"context"
	"data"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Users :
var Users = make(map[string]User)

// InputUser :
type InputUser struct {
}

// User :
type User struct {
	UserID   string `json:"userid" bson:"_id"`
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
	// userID, _ := uuid.NewV4()
	// u.UserID = userID.String()
	// u.Status = "Y"
	// Users[userID.String()] = u
	collection := data.RescueDB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, bson.M{
		"username": u.UserName,
		"email":    u.Email,
		"password": u.Password,
		"usertype": u.UserType,
		"status":   "Y",
	})
	if err != nil {
		return err
	}
	// id := res.InsertedID
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

	collection := data.RescueDB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var users []User
	for cur.Next(ctx) {
		var user User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		return nil, err
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
