package werror

import (
	"errors"
)

type wrapError struct {
	error
	next  error
	frame errors.Frame
}

func (e *wrapError) Unwrap() error {
	return e.next
}

func (e *wrapError) FormatError(p errors.Printer) (next error) {
	p.Print(e.error.Error())
	e.frame.Format(p)
	return e.next
}
