package database

import (
	"context"
	"fmt"
	"log"

	"github.com/hyperxpizza/contact-manager/server/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckIfUsernameTaken(username string) error {
	cursor, err := database.Collection("users").Find(context.Background(), bson.M{"username": username})
	if err != nil {
		log.Printf("Find failed: %v\n", err)
	}

	if cursor.RemainingBatchLength() > 0 {
		return fmt.Errorf("User with this username already exists in the database")
	}

	return nil
}

func CheckIfEmailTaken(email string) error {
	cursor, err := database.Collection("users").Find(context.Background(), bson.M{"email": email})
	if err != nil {
		log.Printf("Find failed: %v\n", err)
	}

	if cursor.RemainingBatchLength() > 0 {
		return fmt.Errorf("User with this email already exists in the database")
	}

	return nil
}

func InsertUser(user model.User) (string, error) {
	res, err := database.Collection("users").InsertOne(context.Background(), &user)
	if err != nil {
		return "", err
	}

	id := fmt.Sprintf("%v", res.InsertedID)
	return id, nil

}
