package gographql

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type ExtendedError struct {
	Err        error
	StatusCode int
}

func (e ExtendedError) Error() string {
	if e.StatusCode >= 500 {
		logrus.Error(e.Err)
		return errors.New("internal server error").Error()
	}
	return errors.Cause(e.Err).Error()
}

func (e ExtendedError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code": e.StatusCode,
	}
}
