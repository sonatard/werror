package werror

import (
	"fmt"

	"golang.org/x/xerrors"
)

type WrapError interface {
	fmt.Formatter
	Unwrap() error
}

func Wrap(err error, next error, calldepth int) WrapError {
	return &wrapError{
		error: err,
		next:  next,
		frame: xerrors.Caller(calldepth),
	}
}
