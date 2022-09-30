package http

import "encoding/json"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Response) JSON() []byte {
	result, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return result
}

func NewResponse(status, message string, data interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
