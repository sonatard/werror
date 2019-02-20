package werror

import (
	"errors"
)

type WrapError interface {
	errors.Wrapper
}

func Wrap(err error, next error, calldepth int) WrapError {
	return &wrapError{
		error: err,
		next:  next,
		frame: errors.Caller(calldepth),
	}
}
