package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/hyperxpizza/contact-manager/server/database"
	"github.com/hyperxpizza/contact-manager/server/graph/generated"
	"github.com/hyperxpizza/contact-manager/server/graph/model"
)

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
