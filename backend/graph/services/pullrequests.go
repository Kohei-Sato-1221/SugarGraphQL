package services

import (
	"context"
	"fmt"
	"log"

	"github.com/Kohei-Sato-1221/SugarGraphQL/backend/generated/model"
	"github.com/Kohei-Sato-1221/SugarGraphQL/backend/graph/db"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type prService struct {
	exec boil.ContextExecutor
}

func (p *prService) GetPullRequestByID(ctx context.Context, id string) (*model.PullRequest, error) {
	pr, err := db.FindPullrequest(ctx, p.exec, id,
		db.PullrequestColumns.ID,
		db.PullrequestColumns.BaseRefName,
		db.PullrequestColumns.Closed,
		db.PullrequestColumns.HeadRefName,
		db.PullrequestColumns.URL,
		db.PullrequestColumns.Number,
		db.PullrequestColumns.Repository,
	)
	if err != nil {
		return nil, err
	}
	return convertPullRequest(pr), nil
}

func convertPullRequestConnection(pullRequests db.PullrequestSlice, hasPrevPage, hasNextPage bool) *model.PullRequestConnection {
	var result model.PullRequestConnection

	for _, dbpr := range pullRequests {
		pr := convertPullRequest(dbpr)

		result.Edges = append(result.Edges, &model.PullRequestEdge{Cursor: pr.ID, Node: pr})
		result.Nodes = append(result.Nodes, pr)
	}
	result.TotalCount = len(pullRequests)

	result.PageInfo = &model.PageInfo{}
	if result.TotalCount != 0 {
		result.PageInfo.StartCursor = &result.Nodes[0].ID
		result.PageInfo.EndCursor = &result.Nodes[result.TotalCount-1].ID
	}
	result.PageInfo.HasPreviousPage = hasPrevPage
	result.PageInfo.HasNextPage = hasNextPage

	return &result
}

func convertPullRequest(pr *db.Pullrequest) *model.PullRequest {
	prURL, err := model.UnmarshalURI(pr.URL)
	if err != nil {
		log.Println("invalid URI", pr.URL)
	}

	return &model.PullRequest{
		ID:          pr.ID,
		BaseRefName: pr.BaseRefName,
		Closed:      (pr.Closed == 1),
		HeadRefName: pr.HeadRefName,
		URL:         prURL,
		Number:      int(pr.Number),
		Repository:  &model.Repository{ID: pr.Repository},
	}
}

func (p *prService) ListPullRequestInRepository(ctx context.Context, repoID, after, before string, first, last int) (*model.PullRequestConnection, error) {
	cond := []qm.QueryMod{
		qm.Select(
			db.PullrequestColumns.ID,
			db.PullrequestColumns.BaseRefName,
			db.PullrequestColumns.Closed,
			db.PullrequestColumns.HeadRefName,
			db.PullrequestColumns.URL,
			db.PullrequestColumns.Number,
			db.PullrequestColumns.Repository,
		),
		db.PullrequestWhere.Repository.EQ(repoID),
	}
	var scanDesc bool

	cond = append(cond,
		qm.OrderBy(fmt.Sprintf("%s asc", db.PullrequestColumns.ID)),
	)

	pullRequests, err := db.Pullrequests(cond...).All(ctx, p.exec)
	if err != nil {
		return nil, err
	}

	var hasNextPage, hasPrevPage bool
	if len(pullRequests) != 0 {
		if scanDesc {
			for i, j := 0, len(pullRequests)-1; i < j; i, j = i+1, j-1 {
				pullRequests[i], pullRequests[j] = pullRequests[j], pullRequests[i]
			}
		}
		startCursor, endCursor := pullRequests[0].ID, pullRequests[len(pullRequests)-1].ID

		var err error
		hasPrevPage, err = db.Pullrequests(
			db.PullrequestWhere.Repository.EQ(repoID),
			db.PullrequestWhere.ID.LT(startCursor),
		).Exists(ctx, p.exec)
		if err != nil {
			return nil, err
		}
		hasNextPage, err = db.Pullrequests(
			db.PullrequestWhere.Repository.EQ(repoID),
			db.PullrequestWhere.ID.GT(endCursor),
		).Exists(ctx, p.exec)
		if err != nil {
			return nil, err
		}
	}

	return convertPullRequestConnection(pullRequests, hasPrevPage, hasNextPage), nil
}
