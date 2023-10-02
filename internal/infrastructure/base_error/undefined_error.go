package base_error

import (
	"github.com/michael-bc/list-full-record-backend/internal/common/common_domain"
)

const UndefinedErrorType = "UndefinedError"

type UndefinedError struct {
	*ErrorData
}

func NewUndefinedError(err error) *UndefinedError {
	return &UndefinedError{
		ErrorData: NewErrorData(common_domain.CommonDomain, UndefinedErrorType, err, nil),
	}
}
