package common

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type DbParameters struct {
	Url             string
	Username        string
	Password        string
	DbName          string
	DroneCollection string
}

func makeURL(parameters DbParameters) string {
	return "mongodb://" + parameters.Url
}

func ConnectMongo(parameters DbParameters) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(makeURL(parameters)))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetDatabase(client *mongo.Client, dbname string) *mongo.Database {
	return client.Database(dbname)
}
