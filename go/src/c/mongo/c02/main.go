package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        string `json:"id,omitempty" bson:"id"`
	Username  string `json:"username,omitempty" bson:"username"`
	Nickname  string `json:"nickname,omitempty" bson:"nickname"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}

	var items []*User
	findOptions := options.Find()
	collection := client.Database("test").Collection("user")
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var item User
		if err := cur.Decode(&item); err != nil {
			log.Fatal(err)
		}
		items = append(items, &item)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	log.Printf("Get items: %v]", items)
}
