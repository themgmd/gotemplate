package dhttp

import (
	"context"
	"gotemplate/internal/auth/types"
	httplib "gotemplate/pkg/http"
	"net/http"
)

type Service interface {
	InitRegistration(context.Context, types.InitRegistrationRequest) (types.InitRegistrationResponse, error)
	FinishRegistration(context.Context, types.FinishRegistrationRequest) error
	Login(context.Context, types.LoginRequest) error
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h Handler) initRegistration(w http.ResponseWriter, r *http.Request) {
	var request types.InitRegistrationRequest

	if err := httplib.ReadBody(r.Body, &request); err != nil {
		httplib.NewBadRequestResponse(w, err)
		return
	}

	resp, err := h.service.InitRegistration(r.Context(), request)
	if err != nil {
		httplib.NewInternalServerErrorResponse(w, err)
		return
	}

	httplib.NewSuccessResponse(w, resp)
}

func (h Handler) finishRegistration(w http.ResponseWriter, r *http.Request) {
	var request types.FinishRegistrationRequest

	// Get registration identifier
	request.Identifier = r.PathValue("identifier")

	if err := httplib.ReadBody(r.Body, &request); err != nil {
		httplib.NewBadRequestResponse(w, err)
		return
	}

	err := h.service.FinishRegistration(r.Context(), request)
	if err != nil {
		httplib.NewInternalServerErrorResponse(w, err)
		return
	}

	httplib.NewCreatedResponse(w, nil)
}

func (h Handler) login(w http.ResponseWriter, r *http.Request) {
	var request types.LoginRequest

	if err := httplib.ReadBody(r.Body, &request); err != nil {
		httplib.NewBadRequestResponse(w, err)
		return
	}

	err := h.service.Login(r.Context(), request)
	if err != nil {
		httplib.NewInternalServerErrorResponse(w, err)
		return
	}

	httplib.NewSuccessResponse(w, nil)
}
