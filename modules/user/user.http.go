package user

import (
	"github.com/labstack/echo/v4"
	"github.com/nanpipat/go-api-template/core"
)

func NewUserHTTP(e *echo.Echo) {
	user := &UserController{}
	e.GET("/users", core.WithHTTPContext(user.Pagination))
}
