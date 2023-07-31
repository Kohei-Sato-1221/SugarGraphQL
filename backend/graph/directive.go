package graph

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Kohei-Sato-1221/SugarGraphQL/backend/generated"
	"github.com/Kohei-Sato-1221/SugarGraphQL/backend/middlewares/auth"
)

var Directive generated.DirectiveRoot = generated.DirectiveRoot{
	IsAuthenticated: IsAuthenticated,
}

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	// TODO
	if _, ok := auth.GetUserName(ctx); !ok {
		return nil, errors.New("not authenticated")
	}
	return next(ctx)
}
