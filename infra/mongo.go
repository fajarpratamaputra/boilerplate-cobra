package infra

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDatabase struct {
	client *mongo.Client
}

func NewMongoDatabase() (*MongoDatabase, error) {
	md := new(MongoDatabase)
	if _, err := md.openClient(""); err != nil {
		return nil, err
	}

	return md, nil
}

func (md *MongoDatabase) openClient(name string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	md.client = client

	return client, err
}

func (md *MongoDatabase) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return md.client.Ping(ctx, readpref.Primary())
}

func (md *MongoDatabase) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return md.client.Disconnect(ctx)
}

func (md *MongoDatabase) GetCollection(databaseName, name string) *mongo.Collection {
	return md.client.Database(databaseName).Collection(name)
}
