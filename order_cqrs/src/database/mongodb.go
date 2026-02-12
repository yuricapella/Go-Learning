package database

import (
	"context"
	"time"

	"github.com/yuricapella/Go-Learning/order_cqrs/src/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// ConnectMongoDB opens connection with MongoDB database and returns client and database
func ConnectMongoDB() (*mongo.Client, *mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(config.MongoDBConnectionString)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, nil, err
	}

	database := client.Database(config.MongoDBDatabaseName)

	return client, database, nil
}
