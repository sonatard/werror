package werror

import (
	"fmt"

	"golang.org/x/xerrors"
)

type wrapError struct {
	error
	next  error
	frame xerrors.Frame
}

func (e *wrapError) Unwrap() error {
	return e.next
}

func (e *wrapError) Format(s fmt.State, v rune) { xerrors.FormatError(e, s, v) }

func (e *wrapError) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.error.Error())
	e.frame.Format(p)
	return e.next
}
