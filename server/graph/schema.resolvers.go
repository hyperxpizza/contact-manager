package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/goware/emailx"
	"github.com/hyperxpizza/contact-manager/server/database"
	"github.com/hyperxpizza/contact-manager/server/graph/generated"
	"github.com/hyperxpizza/contact-manager/server/graph/model"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, username string, email string, password1 string, password2 string) (*model.User, error) {
	//validate username and email
	err := emailx.Validate(email)
	if err != nil {
		return nil, err
	}
	//check if username and email are already taken
	err = database.CheckIfUsernameTaken(username)
	if err != nil {
		return nil, err
	}

	err = database.CheckIfEmailTaken(email)
	if err != nil {
		return nil, err
	}

	//check if passwords are matching
	if password1 != password2 {
		return nil, fmt.Errorf("Passwords do not match")
	}

	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password1), 10)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Username:  username,
		Password:  string(hashedPassword),
		Email:     email,
		IsAdmin:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	id, err := database.InsertUser(user)
	if err != nil {
		return nil, err
	}

	user.ObjectID = &id

	return &user, nil

}

func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*model.AuthPayload, error) {
	payload := model.AuthPayload{
		Token: "token123",
	}

	return &payload, nil
}

func (r *mutationResolver) CreateContact(ctx context.Context, name1 string, name2 *string, surname *string, email *string, phone *string, website *string, company *string) (*model.Contact, error) {
	contact := model.Contact{
		Name1:     name1,
		Name2:     name2,
		Surname:   surname,
		Email:     email,
		Phone:     phone,
		Website:   website,
		Company:   company,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	id, err := database.InsertContact(contact)
	if err != nil {
		return nil, err
	}

	contact.ObjectID = &id

	return &contact, nil
}

func (r *queryResolver) GetContact(ctx context.Context, filter *model.Filter) (*model.Contact, error) {
	contact, err := database.GetContact(filter)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (r *queryResolver) GetContacts(ctx context.Context, filter *model.Filter) ([]*model.Contact, error) {
	contacts, err := database.GetContacts(filter)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func (r *queryResolver) CountContacts(ctx context.Context, filter *model.Filter) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SearchContacts(ctx context.Context, query string) ([]*model.Contact, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) GetAllContacts(ctx context.Context) ([]*model.Contact, error) {
	panic(fmt.Errorf("not implemented"))
}
