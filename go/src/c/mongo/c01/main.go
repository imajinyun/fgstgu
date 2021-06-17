package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	options := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(options)
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}

	at := time.Now().Local().Unix()
	user := User{
		ID:        "1001",
		Username:  "Pony",
		Nickname:  "Pony.Ma",
		CreatedAt: at,
		UpdatedAt: at,
	}
	collection := client.Database("test").Collection("user")
	res, err := collection.InsertOne(context.TODO(), &user)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("✔️ Inserted id object: %v", res.InsertedID)
	id := res.InsertedID.(primitive.ObjectID).Hex()
	log.Printf("✔️ Get string id: %v", id)
	filter := bson.M{"_id": res.InsertedID}
	doc := collection.FindOne(context.TODO(), filter)
	fmt.Println(doc)
	doc.Decode(&user)
	log.Println(user)

	objectID, err := primitive.ObjectIDFromHex("60cb4e73474c3f4451d6e567")
	if err != nil {
		log.Fatal(err)
	}
	row := collection.FindOne(context.TODO(), bson.M{"_id": objectID})

	var bsonObject bson.M
	if err := row.Decode(&bsonObject); err != nil {
		log.Fatal(err)
	}
	log.Println(bsonObject)
}
