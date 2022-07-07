package configs

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoUrl()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	//	ping the database
	err = client.Ping(ctx, nil)
	fmt.Println("Connect to mongodb")
	return client
}

//client instance
var DB *mongo.Client = ConnectDB()

//getting database connection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("go-mongo").Collection(collectionName)
	return collection
}
