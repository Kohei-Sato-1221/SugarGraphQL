package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"fmt"

	"github.com/Kohei-Sato-1221/SugarGraphQL/backend/generated"
	"github.com/Kohei-Sato-1221/SugarGraphQL/backend/generated/model"
)

// Author is the resolver for the author field.
func (r *issueResolver) Author(ctx context.Context, obj *model.Issue) (*model.User, error) {
	return r.Srv.GetUserByID(ctx, obj.Author.ID)
}

// Repository is the resolver for the repository field.
func (r *issueResolver) Repository(ctx context.Context, obj *model.Issue) (*model.Repository, error) {
	return r.Srv.GetRepoByID(ctx, obj.Repository.ID)
}

// ProjectItems is the resolver for the projectItems field.
func (r *issueResolver) ProjectItems(ctx context.Context, obj *model.Issue, after *string, before *string, first *int, last *int) (*model.ProjectV2ItemConnection, error) {
	return r.Srv.ListProjectItemOwnedByIssue(ctx, obj.ID, after, before, first, last)
}

// AddProjectV2ItemByID is the resolver for the addProjectV2ItemById field.
func (r *mutationResolver) AddProjectV2ItemByID(ctx context.Context, input model.AddProjectV2ItemByIDInput) (*model.AddProjectV2ItemByIDPayload, error) {
	panic(fmt.Errorf("not implemented: AddProjectV2ItemByID - addProjectV2ItemById"))
}

// Owner is the resolver for the owner field.
func (r *projectV2Resolver) Owner(ctx context.Context, obj *model.ProjectV2) (*model.User, error) {
	return r.Srv.GetUserByID(ctx, obj.Owner.ID)
}

// Repository is the resolver for the repository field.
func (r *queryResolver) Repository(ctx context.Context, name string, owner string) (*model.Repository, error) {
	user, err := r.Srv.GetUserByName(ctx, owner)
	if err != nil {
		return nil, err
	}
	return r.Srv.GetRepoByFullName(ctx, user.ID, name)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, name string) (*model.User, error) {
	return r.Srv.GetUserByName(ctx, name)
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Issue is the resolver for the issue field.
func (r *queryResolver) Issue(ctx context.Context, id string) (*model.Issue, error) {
	return r.Srv.GetIssue(ctx, id)
}

// Owner is the resolver for the owner field.
func (r *repositoryResolver) Owner(ctx context.Context, obj *model.Repository) (*model.User, error) {
	return r.Srv.GetUserByID(ctx, obj.Owner.ID)
}

// Issue is the resolver for the issue field.
func (r *repositoryResolver) Issue(ctx context.Context, obj *model.Repository, number int) (*model.Issue, error) {
	return r.Srv.GetIssueByRepoAndNumber(ctx, obj.ID, number)
}

// Issues is the resolver for the issues field.
func (r *repositoryResolver) Issues(ctx context.Context, obj *model.Repository, after *string, before *string, first *int, last *int) (*model.IssueConnection, error) {
	return r.Srv.GetIssues(ctx, obj.ID, *after, *before, *first, *last)
}

// PullRequest is the resolver for the pullRequest field.
func (r *repositoryResolver) PullRequest(ctx context.Context, obj *model.Repository, number int) (*model.PullRequest, error) {
	return r.Srv.GetPullRequestByID(ctx, obj.ID)
}

// PullRequests is the resolver for the pullRequests field.
func (r *repositoryResolver) PullRequests(ctx context.Context, obj *model.Repository, after *string, before *string, first *int, last *int) (*model.PullRequestConnection, error) {
	panic(fmt.Errorf("not implemented: PullRequests - pullRequests"))
}

// Issue returns generated.IssueResolver implementation.
func (r *Resolver) Issue() generated.IssueResolver { return &issueResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// ProjectV2 returns generated.ProjectV2Resolver implementation.
func (r *Resolver) ProjectV2() generated.ProjectV2Resolver { return &projectV2Resolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Repository returns generated.RepositoryResolver implementation.
func (r *Resolver) Repository() generated.RepositoryResolver { return &repositoryResolver{r} }

type issueResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type projectV2Resolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type repositoryResolver struct{ *Resolver }
