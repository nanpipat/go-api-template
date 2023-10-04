package emsgs

import "github.com/nanpipat/go-api-template/core"

func NewCustomError(status int, code string, message string) core.Error {
	return core.Error{
		Status:  status,
		Code:    code,
		Message: message,
	}
}
