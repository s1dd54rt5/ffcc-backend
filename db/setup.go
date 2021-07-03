package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	ctx    context.Context
)

func GetDbCollection(title string) *mongo.Collection {
	return Client.Database("ffcc-db").Collection(title)
}

func InitialiseDb() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	dbUrl := os.Getenv("DBURL")
	clientOptions := options.Client().ApplyURI(dbUrl)

	Client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = Client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}
