package stackerror

import (
	"errors"
	"fmt"

	"golang.org/x/xerrors"

	"sample-go-error/pkg/error/code"
)

type Error struct {
	err     error
	pattern code.ErrorPattern
	msg     string
	frame   xerrors.Frame
}

func New(pattern code.ErrorPattern, msg string) error {
	return newError(nil, pattern, msg)
}

func Wrap(err error, pattern code.ErrorPattern, msg string) error {
	return newError(err, pattern, msg)
}

func newError(err error, pattern code.ErrorPattern, msg string) error {
	return &Error{
		err:     err,
		pattern: pattern,
		msg:     msg,
		frame:   xerrors.Caller(2),
	}
}

func Stack(err error) error {
	var pattern code.ErrorPattern
	var msg string
	if e, ok := As(err); ok {
		pattern = e.pattern
		msg = e.msg
		e.msg = "" // 出力が重複しないように引き継いだメッセージを除去
	}
	return &Error{
		err:     err,
		pattern: pattern,
		msg:     msg,
		frame:   xerrors.Caller(1),
	}
}

func As(err error) (*Error, bool) {
	var e *Error
	if errors.As(err, &e) {
		return e, true
	}
	return nil, false
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.pattern.ErrorCode, e.msg)
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Format(s fmt.State, v rune) {
	xerrors.FormatError(e, s, v)
}

func (e *Error) FormatError(p xerrors.Printer) error {
	p.Print(e.msg)
	e.frame.Format(p)
	return e.err
}

func (e *Error) ErrorPattern() code.ErrorPattern {
	if e != nil {
		return e.pattern
	}
	return code.ErrorPattern{}
}
