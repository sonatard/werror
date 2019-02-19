package werror

import (
	"fmt"

	"golang.org/x/xerrors"
)

type WrapError struct {
	Msg   string
	Err   error
	Frame xerrors.Frame
}

func Wrap(wraperr error, cause error) WrapError {
	var newerr WrapError
	newerr.Msg = wraperr.Error()
	newerr.Frame = xerrors.Caller(2)
	newerr.Err = cause
	return newerr
}

func (e *WrapError) Error() string {
	return e.Msg
}

func (e *WrapError) Unwrap() error {
	return e.Err
}

func (e *WrapError) Format(s fmt.State, v rune) { xerrors.FormatError(e, s, v) }

func (e *WrapError) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.Msg)
	e.Frame.Format(p)
	return e.Err
}
