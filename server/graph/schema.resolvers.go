package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/hyperxpizza/contact-manager/server/graph/generated"
	"github.com/hyperxpizza/contact-manager/server/graph/model"
)

func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*model.AuthPayload, error) {
	payload := model.AuthPayload{
		Token: "token123",
	}

	return &payload, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
