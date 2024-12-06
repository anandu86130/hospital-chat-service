package db

import (
	"context"
	"fmt"
	"time"

	"github.com/anandu86130/hospital-chat-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// func ConnectMongoDB(config *config.Config) (*mongo.Client, error) {
// 	clientOptions := options.Client().ApplyURI(config.DBUrl)

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	mongoclient, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
// 	}

// 	err = mongoclient.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
// 	}

// 	log.Println("MongoDB connection established")
// 	return mongoclient, nil
// }

func ConnectToDatabase(cfg *config.Config) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(cfg.DBUrl)
	mongoclient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	fmt.Println("MongoDB connection established")
	return mongoclient.Database(cfg.DBname), nil

}
