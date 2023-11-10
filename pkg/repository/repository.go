package repository

import (
	"sample-go-error/pkg/error/stackerror"
	"sample-go-error/pkg/infra"
)

type Repository struct {
	database *infra.Database
}

func New(database *infra.Database) *Repository {
	return &Repository{database: database}
}

func (r *Repository) ReturnInternalStackError() error {
	if err := r.database.ReturnInternalStackError(); err != nil {
		return stackerror.Stack(err)
	}
	return nil
}

func (r *Repository) ReturnExternalStackError() error {
	if err := r.database.ReturnExternalStackError(); err != nil {
		return stackerror.Stack(err)
	}
	return nil
}

func (r *Repository) ReturnInternalCallersError() error {
	return r.database.ReturnInternalCallersError()
}

func (r *Repository) ReturnExternalCallersError() error {
	return r.database.ReturnExternalCallersError()
}
