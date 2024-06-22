package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DBInstance() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		fmt.Println("connected succesfully")
		panic(err)
	}
	return client
}

var Global *mongo.Client = DBInstance()

func DatabaseTableCreation(p *mongo.Client, Table string) *mongo.Collection {
	r := p.Database("coder range").Collection(Table)
	return r
}
