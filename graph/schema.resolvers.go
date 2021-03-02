package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/elzerver/graphql-services/graph/generated"
	"github.com/elzerver/graphql-services/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, userInput *model.UserInput) (*model.User, error) {
	// panic(fmt.Errorf("not implemented"))

	name := userInput.Name

	var id string
	id = "1"
	user := &model.User{
		ID:    id,
		Name:  &name,
		Posts: []*model.Post{},
	}

	return user, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, postInput *model.PostInput) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
	// post := &model.PostInput{
	// 	title:  "test-post",
	// 	author: "test-author",
	// }

	// return post, nil
}

func (r *mutationResolver) CreateEdgeSitesWorkspace(ctx context.Context, workspaceid int, userid int) (*model.Workspace, error) {
	// panic(fmt.Errorf("not implemented"))
	wsid := workspaceid
	usid := userid

	workspaceCreated := &model.Workspace{
		WorkspaceID: wsid,
		UserID:      usid,
	}

	return workspaceCreated, nil
}

func (r *queryResolver) Me(ctx context.Context, last *int) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
