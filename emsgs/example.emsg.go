package emsgs

import (
	"net/http"

	"github.com/nanpipat/go-api-template/core"
)

var UserNotFound = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "USERS_NOT_FOUND",
	Message: "Users not found",
}
