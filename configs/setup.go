package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))

	if err != nil {
		log.Fatalln("cant connect to db", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatalln("Can't connect to DB erorr -> ", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatalln("connection timeout erorr -> ", err)
	}

	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("golangAPI").Collection(collectionName)
}
