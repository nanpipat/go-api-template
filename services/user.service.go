package services

import (
	"github.com/nanpipat/go-api-template/core"
	"github.com/nanpipat/go-api-template/helpers"
	"github.com/nanpipat/go-api-template/models"
	"github.com/nanpipat/go-api-template/repo"
	"github.com/nanpipat/go-api-template/views"
)

type IUserService interface {
	Pagination(pageOptions *core.PageOptions) (*repo.Pagination[views.User], core.IError)
}

type userService struct {
	ctx core.IContext
}

func NewUserService(ctx core.IContext) IUserService {
	return &userService{
		ctx: ctx,
	}
}

func (s *userService) Pagination(pageOptions *core.PageOptions) (*repo.Pagination[views.User], core.IError) {
	db := repo.New[models.User](s.ctx)

	users, ierr := db.Pagination(pageOptions)
	if ierr != nil {
		return nil, ierr
	}

	return helpers.PaginationMap(users, func(p models.User) views.User {
		return *views.NewUser(&p)
	}), nil
}
