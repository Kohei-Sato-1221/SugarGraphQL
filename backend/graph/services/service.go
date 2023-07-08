package services

import (
	"context"

	"github.com/Kohei-Sato-1221/SugarGraphQL/backend/generated/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Services interface {
	UserService
}

// sql.DBはboil.ContextExecutorの具象として利用可能
func New(exec boil.ContextExecutor) Services {
	return &services{
		userService: &userService{exec: exec},
	}
}

type services struct {
	*userService
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}
