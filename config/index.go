package config

import (
	"context"
	"fmt"
	"k6/constants"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func init() {
	clientoption := options.Client().ApplyURI(constants.Connectstring)

	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected")
	Collection = client.Database(constants.Dbname).Collection(constants.CollectionName)

	fmt.Println("Collection Connected")
}
