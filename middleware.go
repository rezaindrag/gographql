package gographql

import (
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"net/http"
)

func ErrorMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err == nil {
				return nil
			}

			switch errors.Cause(err).(type) {
			case ExtendedError:
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
}
