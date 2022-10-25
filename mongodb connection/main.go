package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	ctx := context.TODO()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln(err)
	}

	userCollection := client.Database("rezhnn").Collection("users")

	var res bson.M
	err = userCollection.FindOne(ctx, bson.D{{"firstName", "reza"}}).Decode(&res)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)

}
