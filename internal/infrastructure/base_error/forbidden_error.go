package base_error

import (
	"github.com/michael-bc/list-full-record-backend/internal/common/common_domain"
)

const ForbiddenErrorType = "ForbiddenError"

type ForbiddenError struct {
	*ErrorData
}

func NewForbiddenError() *ForbiddenError {
	return &ForbiddenError{
		ErrorData: NewErrorData(common_domain.CommonDomain, ForbiddenErrorType, nil, nil),
	}
}
