package data

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var MySigningKey = []byte("mysupersecuret")

func GenerateToken(userDetails UserClaims) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	StandardClaims := StandardClaims{
		userDetails,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	tokenKey := jwt.NewWithClaims(jwt.SigningMethodHS256, StandardClaims)
	token, err := tokenKey.SignedString(MySigningKey)
	if err != nil {
		fmt.Println(err.Error())
		return token, fmt.Errorf("Unable to generate token")
	}
	return token, nil
}

func SetTokenCookies(w http.ResponseWriter, token string) {
	// Set token on cookies
	expirationTime := time.Now().Add(15 * time.Minute)
	http.SetCookie(w, &http.Cookie{
		Name:    "Authorization",
		Value:   token,
		Expires: expirationTime,
	})
}

func UserRegister(userDetails UserClaims, hashpass []byte) error {
	collection := RescueDB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, bson.M{
		"username": userDetails.Username,
		"email":    userDetails.Email,
		"password": hashpass,
		"usertype": userDetails.UserType,
		"status":   "Y",
	})

	return err
}

// GetUser :
func GetUser(email, userType string) (UserClaims, error) {
	var users UserClaims
	collection := RescueDB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.D{primitive.E{Key: "email", Value: email}, primitive.E{Key: "usertype", Value: userType}}).Decode(&users)
	if err != nil {
		return users, err
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
