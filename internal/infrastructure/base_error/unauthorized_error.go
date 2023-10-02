package base_error

import (
	"github.com/michael-bc/list-full-record-backend/internal/common/common_domain"
)

const UnauthorizedErrorType = "UnauthorizedError"

type UnauthorizedError struct {
	*ErrorData
}

func NewUnauthorizedError(devDetails ...string) *UnauthorizedError {
	return &UnauthorizedError{
		ErrorData: NewErrorData(common_domain.CommonDomain, UnauthorizedErrorType, nil, nil, devDetails...),
	}
}
