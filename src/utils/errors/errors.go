package errors

import (
	"fmt"
	"reflect"
	"strings"

	goErr "github.com/go-errors/errors"
)

// ErrContainer error container
type errContainer struct {
	ext     *goErr.Error
	wrapped error
}

// NewError constructor for Error container
func NewError(format string, a ...interface{}) error {
	var err error
	if a == nil {
		err = fmt.Errorf(format)
	} else {
		err = fmt.Errorf(format, a...)
	}

	return errContainer{
		ext:     goErr.New(err),
		wrapped: err,
	}
}

func (err errContainer) Error() string {
	return err.ext.Error()
}

// ErrWrap wraps an error in an error container
// if err is an error container, its returned as is
func ErrWrap(err error) error {
	_, ok := err.(errContainer)
	if ok {
		return err
	}
	return errContainer{
		ext:     goErr.New(err),
		wrapped: err,
	}
}

// ErrIs compares err to an error and returns true if its equal to a given error
func ErrIs(err error, originalErrorType interface{}) bool {
	errCont, ok := err.(errContainer)
	if ok {
		return errCont.is(originalErrorType)
	}
	if orgErr, ok := originalErrorType.(error); ok {
		if err.Error() == orgErr.Error() {
			return true
		}
	}
	return false
}

func (err errContainer) is(originalErrorType interface{}) bool {
	if reflect.TypeOf(err.wrappedErr()) == reflect.TypeOf(originalErrorType) {
		return true
	}
	return false
}

// ErrStack returns the stack trace related to the error
func ErrStack(err error) string {
	errCont, ok := err.(errContainer)
	if ok {
		stackTrace := errCont.errStack()
		stackTrace = strings.Replace(stackTrace, "\n", " - ", -1)
		return stackTrace
	}
	return "no stack trace info in error"
}

func (err errContainer) errStack() string {
	return err.ext.ErrorStack()
}

// WrappedErr returns the wrapped error in the container
func WrappedErr(err error) error {
	if err == nil {
		return nil
	}
	errCont, ok := err.(errContainer)
	if !ok {
		return err
	}
	return errCont.wrappedErr()
}

func (err errContainer) wrappedErr() error {
	return err.wrapped
}
