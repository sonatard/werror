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

// Here is not pointer receiver.
func (e ApplicationError) Wrap(next error) error {
    // set wrap error
    e.err = next
    // set call stack information
    e.frame = xerrors.Caller(1)
    return &e
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

// Here is not pointer receiver.
func (e ApplicationError) Wrap(next error) error {
	e.WrapError = werror.Wrap(&e, next, 2)
	return &e
}

func (e *ApplicationError) Error() string {
	return fmt.Sprintf("%s: code=%d, msg=%s", e.level, e.code, e.msg)
}
```

Usage of your definition error

```go
package main

import (
    // Before Go 1.13
    werror "github.com/sonatard/werror/xerrors"
    // After Go 1.13
    // "github.com/sonatard/werror"
)



func main() {
    err := func1()
    if err != nil {
    	fmt.Fprintf(os.Stderr, "caught error: %+v\n", err)
    }
}

func func1() {
    var ErrUserNotFound = ApplicationError{
    	code:  101,
    	level: "Error",
    	msg:   "not found",
    }

    err := func2()
    if err != nil {
        return ErrUserNotFound.Wrap(err)
    }
}

func func2()error {
    // Before Go 1.13
    return xerrors.New("func2 error")
    // After Go 1.13
    // return errors.New("func2 error")
}
```
