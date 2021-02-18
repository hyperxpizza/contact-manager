package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/hyperxpizza/contact-manager/server/graph/generated"
	"github.com/hyperxpizza/contact-manager/server/graph/model"
)

func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*model.AuthPayload, error) {
	payload := model.AuthPayload{
		Token: "token123",
	}

	return &payload, nil
}

func (r *mutationResolver) CreateContact(ctx context.Context, name1 string, name2 *string, surname *string, email *string, phone *string, website *string) (*model.Contact, error) {
	contact := model.Contact{
		Name1:   name1,
		Name2:   name2,
		Surname: surname,
		Email:   email,
		Phone:   phone,
		Website: website,
	}

}

func (r *queryResolver) GetAllContacts(ctx context.Context) ([]*model.Contact, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
