package data

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RescueDB :
var RescueDB *mongo.Database

func init() {
	initDatabase()
}

func initDatabase() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	RescueDB = client.Database("rescue_planet")
	return
	// client, err := NewClient(options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()
	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
