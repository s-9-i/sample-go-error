package service

import (
	"sample-go-error/pkg/error/stackerror"
	"sample-go-error/pkg/repository"
)

type Service struct {
	repository *repository.Repository
}

func New(repository *repository.Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) ReturnInternalStackError() error {
	if err := s.repository.ReturnInternalStackError(); err != nil {
		return stackerror.Stack(err)
	}
	return nil
}

func (s *Service) ReturnExternalStackError() error {
	if err := s.repository.ReturnExternalStackError(); err != nil {
		return stackerror.Stack(err)
	}
	return nil
}

func (s *Service) ReturnInternalCallersError() error {
	return s.repository.ReturnInternalCallersError()
}

func (s *Service) ReturnExternalCallersError() error {
	return s.repository.ReturnExternalCallersError()
}
