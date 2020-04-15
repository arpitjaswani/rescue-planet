package models

import (
	"context"
	"data"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
)

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

	return err
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
	if err = cur.All(ctx, &users); err != nil {
		return nil, err
	}

	// for cur.Next(ctx) {
	// 	var user User
	// 	err := cur.Decode(&user)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	users = append(users, user)
	// }

	// if err := cur.Err(); err != nil {
	// 	return nil, err
	// }

	return users, nil
}

// UpdateUser :
func (u User) UpdateUser() (User, error) {
	var user User
	userID, _ := primitive.ObjectIDFromHex(u.UserID)
	// opts := options.Update().SetUpsert(true)
	collection := data.RescueDB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.UpdateOne(ctx, bson.M{
		"_id": userID,
	}, bson.M{
		"$set": bson.M{
			"email":    u.Email,
			"password": u.Password,
		},
	})

	if err != nil {
		return user, err
	}

	if result.ModifiedCount == 0 {
		return user, errors.New("User not found")
	}
	return user, nil
}

// DeactivateUser :
func DeactivateUser(userID string) (User, error) {
	var user User
	id, _ := primitive.ObjectIDFromHex(userID)
	collection := data.RescueDB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.UpdateOne(ctx, bson.M{
		"_id": id,
	}, bson.M{
		"$set": bson.M{
			"status": "N",
		},
	})

	if err != nil {
		return user, err
	}

	if result.ModifiedCount == 0 {
		return user, errors.New("User not found")
	}
	return user, nil
}
