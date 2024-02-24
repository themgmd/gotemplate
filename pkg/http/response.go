package http

import (
	"github.com/goccy/go-json"
	"gotemplate/pkg/customerror"
	"gotemplate/pkg/pagination"
	"log/slog"
	"net/http"
)

type Payload interface{}

type BaseResponse struct {
	Success bool `json:"success"`
}

type ErrorResponse struct {
	BaseResponse
	customerror.Error
}

type ListResponse struct {
	BaseResponse
	Meta pagination.ResponsePagination
	Data Payload `json:"data"`
}

type Response struct {
	BaseResponse
	Data Payload `json:"data,omitempty"`
}

func bytes(resp any) ([]byte, error) {
	return json.Marshal(resp)
}

func NewListResponse(w http.ResponseWriter, pag pagination.ResponsePagination, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := ListResponse{
		BaseResponse: BaseResponse{
			Success: true,
		},
		Meta: pag,
		Data: data,
	}

	respBytes, err := bytes(response)
	if err != nil {
		slog.Error("Error occurred while response marshalling json", err)
		return
	}

	_, err = w.Write(respBytes)
	if err != nil {
		slog.Error("Error occurred while writes response", err)
		return
	}
}

func NewSuccessResponse(w http.ResponseWriter, data interface{}) {
	newResponseWithData(w, http.StatusOK, data)
}

func NewCreatedResponse(w http.ResponseWriter, data interface{}) {
	newResponseWithData(w, http.StatusCreated, data)
}

func NewBadRequestResponse(w http.ResponseWriter, err error) {
	newErrorResponse(w, http.StatusBadRequest, err)
}

func NewInternalServerErrorResponse(w http.ResponseWriter, err error) {
	newErrorResponse(w, http.StatusInternalServerError, err)
}

func newResponseWithData(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
		BaseResponse: BaseResponse{
			Success: true,
		},
		Data: data,
	}

	respBytes, err := bytes(response)
	if err != nil {
		slog.Error("Error occurred while response marshalling json", err)
		return
	}

	_, err = w.Write(respBytes)
	if err != nil {
		slog.Error("Error occurred while writes response", err)
		return
	}
}

func newErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	customErr := customerror.FromError(err)

	response := ErrorResponse{
		BaseResponse: BaseResponse{
			Success: false,
		},
		Error: *customErr,
	}

	respBytes, respErr := bytes(response)
	if respErr != nil {
		slog.Error("Error occurred while response marshalling json:", respErr)
		return
	}

	_, respErr = w.Write(respBytes)
	if respErr != nil {
		slog.Error("Error occurred while writes response:", respErr)
		return
	}
}
