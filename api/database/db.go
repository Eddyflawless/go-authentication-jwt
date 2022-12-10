package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var dbClient *mongo.Client

func CreateDBConnection() {

	mongoDBURL := os.Getenv("MONGODB_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDBURL))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	dbClient = client
}

func OpenCollection(collectionName string) *mongo.Collection {

	database := os.Getenv("MONGO_DB")

	return dbClient.Database(database).Collection(collectionName)

}

func ConvertObjectIDToHex(id string) (objectId primitive.ObjectID, err error) {

	objectId, err = primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Invalid id")
		return objectId, err
	}

	return objectId, nil

}
