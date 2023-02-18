package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	url        = "mongodb://localhost:27017"
	db         = "test"
	collection = "dev"
)

type Client struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal("Connect to mongo error: ", err)
	}
	coll := client.Database(db).Collection(collection)

	distinct, err := coll.Distinct(ctx, "_id", bson.M{})
	if err != nil {
		return
	}
	fmt.Println(distinct)
}

func List(ctx context.Context, collection *mongo.Collection) {
	opts := options.Find()
	filter := bson.M{}
	documents, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return
	}
	fmt.Println(documents)

	filter["_id"] = "id"
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return
	}

	var clients []*Client
	if err = cursor.All(ctx, &clients); err != nil {
		return
	}
}

func DeleteClient(ctx context.Context, collection *mongo.Collection) {
	filter := bson.M{}
	filter["_id"] = "id"
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return
	}

	if result != nil && result.DeletedCount == 0 {
		return
	}
}
