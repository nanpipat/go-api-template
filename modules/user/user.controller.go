package user

import (
	"net/http"

	"github.com/nanpipat/go-api-template/core"
	"github.com/nanpipat/go-api-template/services"
)

type UserController struct {
}

func (m *UserController) Pagination(c core.IHTTPContext) error {
	svc := services.NewUserService(c)

	result, ierr := svc.Pagination(c.GetPageOptions())
	if ierr != nil {
		return c.JSON(ierr.GetStatus(), ierr.JSON())
	}

	return c.JSON(http.StatusOK, result)
}
