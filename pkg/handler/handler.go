package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"sample-go-error/pkg/error/stackerror"
	"sample-go-error/pkg/usecase"
)

type Handler struct {
	usecase *usecase.Usecase
}

func New(usecase *usecase.Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) ReturnInternalStackError(c echo.Context) error {
	if err := h.usecase.ReturnInternalStackError(); err != nil {
		return stackerror.Stack(err)
	}
	return c.String(http.StatusOK, "Hello, World!")
}

func (h *Handler) ReturnExternalStackError(c echo.Context) error {
	if err := h.usecase.ReturnExternalStackError(); err != nil {
		return stackerror.Stack(err)
	}
	return c.String(http.StatusOK, "Hello, World!")
}

func (h *Handler) ReturnInternalCallersError(c echo.Context) error {
	if err := h.usecase.ReturnInternalCallersError(); err != nil {
		return err
	}
	return c.String(http.StatusOK, "Hello, World!")
}

func (h *Handler) ReturnExternalCallersError(c echo.Context) error {
	if err := h.usecase.ReturnExternalCallersError(); err != nil {
		return err
	}
	return c.String(http.StatusOK, "Hello, World!")
}
