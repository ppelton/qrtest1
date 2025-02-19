package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"

	"github.com/google/uuid"
	"qrtest1.peltonium.net/graph/model"
)

// CreateItem is the resolver for the createItem field.
func (r *mutationResolver) CreateItem(ctx context.Context, input model.ItemInput) (*model.Item, error) {
	item := &model.Item{
		ID:   uuid.NewString(),
		Name: input.Name,
	}
	r.items = append(r.items, item)
	return item, nil
}

// Items is the resolver for the items field.
func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	return r.items, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
