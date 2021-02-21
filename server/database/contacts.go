package database

import (
	"context"
	"fmt"
	"log"

	model "github.com/hyperxpizza/contact-manager/server/graph/model"
)

//InsertContact created a new contact document
func InsertContact(contact model.Contact) (string, error) {
	res, err := database.Collection("contacts").InsertOne(context.Background(), &contact)
	if err != nil {
		log.Printf("Error while InsertContract: %v", err)
		return "", err
	}

	id := fmt.Sprintf("%v", res.InsertedID)

	return id, nil
}

//GetContact returns a single contact based on filter
func GetContact(filter *model.Filter) (*model.Contact, error) {
	var contact model.Contact
	err := database.Collection("contacts").FindOne(context.Background(), filter).Decode(&contact)
	if err != nil {
		log.Printf("GetContact failed: %v\n", err)
		return nil, err
	}

	return &contact, nil
}

//GetContacts returns multiple contacts based on filter
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

func CountContacts(filter *model.Filter) (int, error) {
	count, err := database.Collection("contacts").CountDocuments(context.Background(), filter)
	if err != nil {
		log.Printf("CountContacts failed: %v\n", err)
		return 0, err
	}

	return int(count), nil
}
