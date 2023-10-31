package http

import (
	"github.com/goccy/go-json"
	"gotemplate/pkg/pagination"
	"log/slog"
	"net/http"
)

type Payload interface{}

type BaseResponse struct {
	Success bool   `json:"success"`
	Comment string `json:"comment,omitempty"`
}

type ListResponse struct {
	BaseResponse
	Meta pagination.ResponsePagination
	Data Payload `json:"data"`
}

func (lr ListResponse) Bytes() ([]byte, error) {
	return json.Marshal(lr)
}

type Response struct {
	BaseResponse
	Data Payload `json:"data"`
}

func (r Response) Bytes() ([]byte, error) {
	return json.Marshal(r)
}

func NewSuccessResponse(w http.ResponseWriter, statusCode StatusCode, data interface{}) {
	if statusCode.IsServerError() || statusCode.IsClientError() {
		slog.Error("status code must me <= 400")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode.Int())

	response := Response{
		BaseResponse: BaseResponse{
			Success: true,
		},
		Data: data,
	}

	respBytes, err := response.Bytes()
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

func NewListResponse(w http.ResponseWriter, pag pagination.ResponsePagination, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(StatusOK.Int())

	response := ListResponse{
		BaseResponse: BaseResponse{
			Success: true,
		},
		Meta: pag,
		Data: data,
	}

	respBytes, err := response.Bytes()
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

func NewErrorResponse(w http.ResponseWriter, statusCode StatusCode, err error) {
	if statusCode.IsInfo() || statusCode.IsSuccess() {
		slog.Error("status code must me greater or equals 400")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode.Int())

	response := Response{
		BaseResponse: BaseResponse{
			Success: false,
			Comment: err.Error(),
		},
	}

	respBytes, err := response.Bytes()
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
