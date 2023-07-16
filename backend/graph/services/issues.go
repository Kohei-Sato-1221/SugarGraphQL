package services

import (
	"context"
	"log"

	"github.com/Kohei-Sato-1221/SugarGraphQL/backend/generated/model"
	"github.com/Kohei-Sato-1221/SugarGraphQL/backend/graph/db"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type issueService struct {
	exec boil.ContextExecutor
}

func (i *issueService) GetIssueByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.Issue, error) {
	issue, err := db.Issues(
		qm.Select(
			db.IssueColumns.ID,
			db.IssueColumns.URL,
			db.IssueColumns.Title,
			db.IssueColumns.Closed,
			db.IssueColumns.Number,
			db.IssueColumns.Author,
			db.IssueColumns.Repository,
		),
		db.IssueWhere.Repository.EQ(repoID),
		db.IssueWhere.Number.EQ(number),
	).One(ctx, i.exec)
	if err != nil {
		return nil, err
	}
	return convertIssue(issue), nil
}

func (i *issueService) GetIssues(ctx context.Context, repoID string, after, before string, first, last int) (*model.IssueConnection, error) {
	cond := []qm.QueryMod{
		qm.Select(
			db.IssueColumns.ID,
			db.IssueColumns.URL,
			db.IssueColumns.Title,
			db.IssueColumns.Closed,
			db.IssueColumns.Number,
			db.IssueColumns.Author,
			db.IssueColumns.Repository,
		),
		db.IssueWhere.Repository.EQ(repoID),
	}
	issues, err := db.Issues(cond...).All(ctx, i.exec)
	if err != nil {
		return nil, err
	}
	return convertIssueConnection(issues, false, false), nil
	// return convertIssueConnection(issues, hasPrevPage, hasNextPage), nil
}

func (i *issueService) GetIssue(ctx context.Context, id string) (*model.Issue, error) {
	issue, err := db.Issues(
		qm.Select(
			db.IssueColumns.ID,
			db.IssueColumns.URL,
			db.IssueColumns.Title,
			db.IssueColumns.Closed,
			db.IssueColumns.Number,
			db.IssueColumns.Author,
			db.IssueColumns.Repository,
		),
		db.IssueWhere.ID.EQ(id),
	).One(ctx, i.exec)
	if err != nil {
		return nil, err
	}
	return convertIssue(issue), nil
}

func convertIssue(issue *db.Issue) *model.Issue {
	issueURL, err := model.UnmarshalURI(issue.URL)
	if err != nil {
		log.Println("invalid URI", issue.URL)
	}

	return &model.Issue{
		ID:         issue.ID,
		URL:        issueURL,
		Title:      issue.Title,
		Closed:     (issue.Closed == 1),
		Number:     int(issue.Number),
		Author:     &model.User{ID: issue.Author},
		Repository: &model.Repository{ID: issue.Repository},
	}
}

func convertIssueConnection(issues db.IssueSlice, hasPrevPage, hasNextPage bool) *model.IssueConnection {
	var result model.IssueConnection

	for _, dbi := range issues {
		issue := convertIssue(dbi)

		result.Edges = append(result.Edges, &model.IssueEdge{Cursor: issue.ID, Node: issue})
		result.Nodes = append(result.Nodes, issue)
	}
	result.TotalCount = len(issues)

	result.PageInfo = &model.PageInfo{}
	if result.TotalCount != 0 {
		result.PageInfo.StartCursor = &result.Nodes[0].ID
		result.PageInfo.EndCursor = &result.Nodes[result.TotalCount-1].ID
	}
	result.PageInfo.HasPreviousPage = hasPrevPage
	result.PageInfo.HasNextPage = hasNextPage

	return &result
}
