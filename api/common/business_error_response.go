package common

import (
	"RESTfulMS1/business"
	"net/http"
)

var (
	ErrInternalServer string = "500"
	ErrDataConflict   string = "409"
	ErrDataNotSpec    string = "400"
)

type BusinessErrorResponseSpec struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewBusinessErrorResponse(err error) (int, *BusinessErrorResponseSpec) {
	switch err {
	default:
		return errResponseInternalServerError()

	case business.ErrDataConflict:
		return errResponseDataConflict(err.Error())

	case business.ErrDataNotSpec:
		return errResponseDataNotSpec(err.Error())

	}
}

func errResponseInternalServerError() (int, *BusinessErrorResponseSpec) {
	return http.StatusInternalServerError, &BusinessErrorResponseSpec{
		Code:    ErrInternalServer,
		Message: "Internal Server Error",
	}
}

func errResponseDataConflict(message string) (int, *BusinessErrorResponseSpec) {
	return http.StatusConflict, &BusinessErrorResponseSpec{
		Code:    ErrDataConflict,
		Message: message,
	}
}

func errResponseDataNotSpec(message string) (int, *BusinessErrorResponseSpec) {
	return http.StatusBadRequest, &BusinessErrorResponseSpec{
		Code:    ErrDataNotSpec,
		Message: message,
	}
}
