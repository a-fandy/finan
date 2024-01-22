package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(config Config) (*mongo.Client, error) {
	// Replace with your MongoDB connection string
	username := config.Get("MONGO_USERNAME")
	password := config.Get("MONGO_PASSWORD")
	host := config.Get("MONGO_HOST")
	port := config.Get("MONGO_PORT")
	dbName := config.Get("MONGO_DB_NAME")
	connectionString := "mongodb://" + username + ":" + password + "@" + host + ":" + port + "/?authSource=" + dbName
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
