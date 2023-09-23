package transport

import "encoding/json"

type BaseResponse struct {
	Success bool   `json:"success"`
	Comment string `json:"comment"`
}

func NewBaseResponse(success bool, comment string) *BaseResponse {
	return &BaseResponse{success, comment}
}

func (br BaseResponse) Bytes() []byte {
	bt, _ := json.Marshal(br)
	return bt
}

type Response struct {
	BaseResponse
	Data interface{} `json:"data"`
}

func (r Response) Bytes() []byte {
	bt, _ := json.Marshal(r)
	return bt
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		BaseResponse: *NewBaseResponse(true, ""),
		Data:         data,
	}
}

func NewErrorResponse(comment string) *Response {
	return &Response{
		BaseResponse: *NewBaseResponse(false, comment),
		Data:         nil,
	}
}
