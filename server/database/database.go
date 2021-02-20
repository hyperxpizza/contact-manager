package database

import (
	"context"
	"fmt"
	"log"
	"time"

	model "github.com/hyperxpizza/contact-manager/server/graph/model"
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

func InsertContact(contact model.Contact) (string, error) {
	res, err := database.Collection("contacts").InsertOne(context.Background(), &contact)
	if err != nil {
		log.Printf("Error while InsertContract: %v", err)
		return "", err
	}

	id := fmt.Sprintf("%v", res.InsertedID)

	return id, nil
}

func GetContact(filter *model.Filter) (*model.Contact, error) {
	var contact model.Contact
	err := database.Collection("contacts").FindOne(context.Background(), filter).Decode(&contact)
	if err != nil {
		log.Printf("GetContact failed: %v\n", err)
		return nil, err
	}

	return &contact, nil
}

func GetContacts(filter *model.Filter) ([]*model.Contact, error) {
	var contacts []*model.Contact

	cursor, err := database.Collection("contacts").Find(context.Background(), filter)
	if err != nil {
		log.Printf("GetContacts failed: %v\n", err)
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var contact model.Contact

		err := cursor.Decode(&contact)
		if err != nil {
			log.Printf("Cursor.Decode failed: %v\n", err)
			return nil, err
		}

		contacts = append(contacts, &contact)
	}

	return contacts, nil
}
