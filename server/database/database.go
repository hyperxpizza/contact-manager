package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var database *mongo.Database

func ConnectToDB(username, password, authSource string, host, port int) (*mongo.Client, context.Context, context.CancelFunc) {

}

func InsertContact(model.Contact)
