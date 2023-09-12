package services

import (
	"context"
	"fmt"
	"k6/config"
	"k6/models"
	"log"
)

func Insert(profile models.RequestData) {
	inserted, err := config.Collection.InsertOne(context.Background(), profile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted", inserted.InsertedID)
}
