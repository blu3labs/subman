package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var db *mongo.Database
var subPayments *mongo.Collection
var ctx = context.Background()

func init() {
	Connect()
	SetUp()
}

func Connect() {
	mongoURI, success := os.LookupEnv("MONGO_URI")
	if !success {
		panic("MONGO_URI not found")
	}
	options := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		panic(err)
	}
	db = client.Database("subman")
	subPayments = db.Collection("subPayments")
}

func SetUp() {
}
