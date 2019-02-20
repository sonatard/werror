package werror_test

import (
	"fmt"
	"os"

	werror "github.com/sonatard/werror/xerrors"
	"golang.org/x/xerrors"
)

type ApplicationError struct {
	level string
	code  int
	msg   string
	werror.WrapError
}

func (e *ApplicationError) Wrap(wraperr error) error {
	newerr := *e
	newerr.WrapError = *werror.Wrap(e, wraperr, 2)
	return &newerr
}

func (e *ApplicationError) Error() string {
	return fmt.Sprintf("%s: code=%d, msg=%s", e.level, e.code, e.msg)
}

func Example() {
	// NOTE: We define ApplicationError like the comment below.
	// (We show this as a comment since we cannot define methods in a function)
	//
	// type ApplicationError struct {
	//     level string
	//     code  int
	//     msg   string
	//     werror.WrapError
	// }
	//
	// func (e *ApplicationError) Wrap(wraperr error) error {
	//     newerr := *e
	//     newerr.WrapError = werror.Wrap(e, wraperr)
	//     return &newerr
	// }
	//
	// func (e *ApplicationError) Error() string {
	//     return fmt.Sprintf("%s: code=%d, msg=%s", e.level, e.code, e.msg)
	// }

	var ErrUserNotFound = &ApplicationError{
		code:  101,
		level: "Error",
		msg:   "not found",
	}

	a := func() error {
		return xerrors.New("error in a")
	}

	err := ErrUserNotFound.Wrap(a())
	fmt.Fprintf(os.Stderr, "caught error: %+v\n", err)

	// NOTE: The some output like below will be printed to stderr,
	// but we cannot use this as expected test result since the full path
	// will be different when this example runs in a different directory.
	//
	// caught error: Error: code=101, msg=not found:
	//    github.com/sonatard/werror_test.Example
	//        /tmp/go/src/github.com/sonatard/werror/example_test.go:59
	//  - error in a:
	//    github.com/sonatard/werror_test.Example.func1
	//        /tmp/go/src/github.com/sonatard/werror/example_test.go:56

	// NOTE: Below "Output:" is needed to run this example as a test by
	// go test -v

	// Output:
}
