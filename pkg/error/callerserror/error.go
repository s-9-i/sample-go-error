package callerserror

import (
	"errors"
	"fmt"
	"io"
	"runtime"
	"strings"

	"sample-go-error/pkg/error/code"
)

type Error struct {
	err     error
	pattern code.ErrorPattern
	msg     string
	frames  *frames
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
		frames:  callers(),
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
	switch v {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%s\n%s", e.Error(), strings.Join(e.frames.GetStackTrace(), "\n"))
			if e.Unwrap() != nil {
				_, _ = fmt.Fprintf(s, "\n- %+v", e.Unwrap())
			}
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", e.Error())
	}
}

func (e *Error) ErrorPattern() code.ErrorPattern {
	if e != nil {
		return e.pattern
	}
	return code.ErrorPattern{}
}

// https://github.com/getsentry/sentry-go/blob/8f8897dc8b964b6116f737c6e78ecd55af0f90dd/stacktrace.go#L84
func (e *Error) StackTrace() frames {
	if e.frames == nil {
		return nil
	}
	return *e.frames
}

type frames []uintptr

func (f *frames) GetStackTrace() []string {
	if f == nil {
		return nil
	}
	fs := runtime.CallersFrames(*f)
	stackTrace := make([]string, 0)
	for {
		frame, more := fs.Next()
		stackTrace = append(stackTrace, fmt.Sprintf("\t%s\n\t\t%s:%d", frame.Function, frame.File, frame.Line))
		if !more {
			break
		}
	}
	return stackTrace
}

func callers() *frames {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(4, pcs[:])
	var st frames = pcs[0:n]
	return &st
}
