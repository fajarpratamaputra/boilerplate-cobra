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

func NewMongoDatabase(ctx context.Context) (*MongoDatabase, error) {
	md := new(MongoDatabase)
	if _, err := md.openClient(ctx, "mongodb://localhost:27017"); err != nil {
		return nil, err
	}

	return md, nil
}

func (md *MongoDatabase) openClient(ctx context.Context, dbUrl string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))

	md.client = client

	return client, err
}

func (md *MongoDatabase) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	return md.client.Ping(ctx, readpref.Primary())
}

func (md *MongoDatabase) Close(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return md.client.Disconnect(ctx)
}

func (md *MongoDatabase) GetCollection(databaseName, name string) *mongo.Collection {
	return md.client.Database(databaseName).Collection(name)
}
