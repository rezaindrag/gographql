package gographql

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"net/http"
)

func ErrorMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err == nil {
				fmt.Println(2)
				return nil
			}
			fmt.Println(1)
			switch errors.Cause(err).(type) {
			case ExtendedError:
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
}
