package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DBInstance func connect the database.
func DBInstance() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		fmt.Println("connected successfully")
		panic(err)
	}
	return client
}

var Global *mongo.Client = DBInstance()

// DatabaseTableCreation func create the database.
func DatabaseTableCreation(p *mongo.Client, table string) *mongo.Collection {
	r := p.Database("restaurant").Collection(table)
	return r
}
