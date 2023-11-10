package infra

import (
	"strconv"

	"sample-go-error/pkg/error/callerserror"
	"sample-go-error/pkg/error/code"
	"sample-go-error/pkg/error/stackerror"
)

type Database struct {
}

func New() *Database {
	return &Database{}
}

func (d *Database) ReturnInternalStackError() error {
	return stackerror.New(code.InternalServerError, "エラーが発生しました")
}

func (d *Database) ReturnExternalStackError() error {
	_, err := strconv.Atoi("dummy")
	return stackerror.Wrap(err, code.InternalServerError, "エラーが発生しました")
}

func (d *Database) ReturnInternalCallersError() error {
	return callerserror.New(code.InternalServerError, "エラーが発生しました")
}

func (d *Database) ReturnExternalCallersError() error {
	_, err := strconv.Atoi("dummy")
	return callerserror.Wrap(err, code.InternalServerError, "エラーが発生しました")
}
