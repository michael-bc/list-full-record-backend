package http_server

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/michael-bc/list-full-record-backend/internal/common/common_domain"
	"github.com/michael-bc/list-full-record-backend/internal/infrastructure/base_error"
	"github.com/michael-bc/list-full-record-backend/internal/infrastructure/configs"
	"github.com/michael-bc/list-full-record-backend/internal/infrastructure/logger"
	"net/http"
)

func ErrorHandler(log logger.Logger, environment configs.Environment) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		if httpError, ok := err.(*echo.HTTPError); ok {
			if httpError.Internal != nil {
				if internalHttpErr, ok := httpError.Internal.(*echo.HTTPError); ok {
					httpError = internalHttpErr
				}
			}
			if httpError.Code == 404 {
				err = base_error.NewNotFoundError(common_domain.CommonDomain)
			}
		}

		var statusCode int

		switch err.(type) {
		case *base_error.BadRequestError,
			*base_error.ValidationError:
			statusCode = http.StatusBadRequest

		case *base_error.UnauthorizedError:
			statusCode = http.StatusUnauthorized

		case *base_error.ForbiddenError:
			statusCode = http.StatusForbidden

		case *base_error.NotFoundError:
			statusCode = http.StatusNotFound

		case *base_error.DatabaseError:
			statusCode = http.StatusInternalServerError

		default:
			statusCode = http.StatusInternalServerError
			err = base_error.NewUndefinedError(err)
		}

		baseError := err.(base_error.BaseError)

		errorData := baseError.GetErrorData()
		errorData.Data["path"] = fmt.Sprintf("%s %s", c.Request().Method, c.Request().URL)

		jsonErr, _ := json.Marshal(err)
		if statusCode < 500 {
			log.Debug(string(jsonErr))
		} else {
			log.Error(string(jsonErr))
		}

		if environment != configs.DevelopmentEnvironment {
			c.JSON(statusCode, base_error.TruncateErrorData(errorData))
		} else {
			c.JSON(statusCode, baseError)
		}
	}
}
