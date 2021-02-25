package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/elzerver/graphql-services/graph/generated"
	"github.com/elzerver/graphql-services/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, name string) (*model.User, error) {
	// panic(fmt.Errorf("not implemented"))
	var id string
	id = "1"
	user := &model.User{
		ID:   id,
		Name: &name,
	}
	r.user = append(r.user, user)
	return user, nil
}

func (r *queryResolver) Me(ctx context.Context) ([]*model.User, error) {
	// panic(fmt.Errorf("not implemented"))
	return r.user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
