package usecase

import (
	"sample-go-error/pkg/error/stackerror"
	"sample-go-error/pkg/service"
)

type Usecase struct {
	service *service.Service
}

func New(service *service.Service) *Usecase {
	return &Usecase{service: service}
}

func (u *Usecase) ReturnInternalStackError() error {
	if err := u.service.ReturnInternalStackError(); err != nil {
		return stackerror.Stack(err)
	}
	return nil
}

func (u *Usecase) ReturnExternalStackError() error {
	if err := u.service.ReturnExternalStackError(); err != nil {
		return stackerror.Stack(err)
	}
	return nil
}

func (u *Usecase) ReturnInternalCallersError() error {
	return u.service.ReturnInternalCallersError()
}

func (u *Usecase) ReturnExternalCallersError() error {
	return u.service.ReturnExternalCallersError()
}
