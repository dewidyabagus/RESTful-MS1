package common

import "net/http"

var (
	ErrBadRequest string = "400"
)

type ControllerErrorResponseSpec struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func BadRequestResponse() (int, *ControllerErrorResponseSpec) {
	return http.StatusBadRequest, &ControllerErrorResponseSpec{
		Code:    ErrBadRequest,
		Message: "Bad Request",
	}
}
