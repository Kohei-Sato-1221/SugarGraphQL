package services

import (
	"context"

	"github.com/Kohei-Sato-1221/SugarGraphQL/backend/generated/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Services interface {
	UserService
	RepositoryService
	IssueService
	PrService
	ProjectItemService
}

// sql.DBはboil.ContextExecutorの具象として利用可能
func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:        &userService{exec: exec},
		repositoryService:  &repositoryService{exec: exec},
		issueService:       &issueService{exec: exec},
		prService:          &prService{exec: exec},
		projectItemService: &projectItemService{exec: exec},
	}
}

type services struct {
	*userService
	*repositoryService
	*issueService
	*prService
	*projectItemService
}

type UserService interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type RepositoryService interface {
	GetRepoByID(ctx context.Context, id string) (*model.Repository, error)
	GetRepoByFullName(ctx context.Context, owner, name string) (*model.Repository, error)
}

type IssueService interface {
	GetIssue(ctx context.Context, id string) (*model.Issue, error)
	GetIssueByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.Issue, error)
	GetIssues(ctx context.Context, repoID string, after, before string, first, last int) (*model.IssueConnection, error)
}

type PrService interface {
	GetPullRequestByID(ctx context.Context, id string) (*model.PullRequest, error)
	ListPullRequestInRepository(ctx context.Context, repoID, after, before string, first, last int) (*model.PullRequestConnection, error)
}

type ProjectItemService interface {
	ListProjectItemOwnedByIssue(ctx context.Context, issueID string, after *string, before *string, first *int, last *int) (*model.ProjectV2ItemConnection, error)
}
