package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s/?authSource=%s"
)

var database *mongo.Database
var mongoCtx context.Context
var cancelFunc context.CancelFunc

//ConnectToDB creates a connection to the mongo database and sets global client and context
func ConnectToDB(username, password, authSource, host string, port int) {
	clusterEndpoint := fmt.Sprintf("%s:%d", host, port)

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint, authSource)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	database = client.Database("manager")
	mongoCtx = ctx
	cancelFunc = cancel

	fmt.Println("Connected to MongoDB!")
}
