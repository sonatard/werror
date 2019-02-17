# werror

`werror` provides Wrap error to your definition error type.


## Sample

Not use `werror`

```go
type ApplicationError struct {
    level string
    code  int
    msg   string
    // wrap target error
    err   error
    // call stack information
    frame xerrors.Frame
}

func (e *ApplicationError) Error() string {
    return fmt.Sprintf("%s: code=%d, msg=%s", e.level, e.code, e.msg)
}

func (e *ApplicationError) Unwrap() error {
    return e.err
}

// implement fmt.Formatter
func (e *ApplicationError) Format(s fmt.State, v rune) { xerrors.FormatError(e, s, v) }

// implement xerrors.Formatter
func (e *ApplicationError) FormatError(p xerrors.Printer) (next error) {
    p.Print(e.Error())
    e.frame.Format(p)
    return e.err
}

func ApplicationErrorWithWrap(err *ApplicationError, wraperr error) error {
    newerr := *err
    // set call stack information
    newerr.frame = xerrors.Caller(1)
    // set wrap error
    newerr.err = wraperr
    return &newerr
}

```

Use `werror`

```go
type ApplicationError struct {
	level string
	code  int
	msg   string
	// embededd
	werror.WrapError
}

func (e *ApplicationError) Error() string {
	return fmt.Sprintf("%s: code=%d, msg=%s", e.level, e.code, e.msg)
}

func ApplicationErrorWithWrap(err *ApplicationError, wraperr error) error {
	newerr := *err
	newerr.WrapError.Msg = err.Error()
	newerr.WrapError.Frame = xerrors.Caller(1)
	newerr.WrapError.Err = wraperr
	return &newerr
}

```
