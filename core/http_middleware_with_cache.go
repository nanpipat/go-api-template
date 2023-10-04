package core

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HTTPWithCache(key func(IHTTPContext) string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := c.(IHTTPContext)
			var item interface{}
			err := cc.Cache().GetJSON(&item, key(cc))
			if err != nil && !errors.Is(err, redis.Nil) {
				cc.NewError(err, Error{
					Status:  http.StatusInternalServerError,
					Code:    "DATABASE_ERROR",
					Message: "database internal error"})
			}

			if item != nil {
				return c.JSON(http.StatusOK, item)
			}

			return next(c)
		}
	}
}
