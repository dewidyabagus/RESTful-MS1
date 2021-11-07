package common

import "net/http"

var (
	Success        string = "200"
	SuccessCreated string = "201"
)

type SuccessResponseSpec struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponseWithoutData() (int, *SuccessResponseSpec) {
	return http.StatusOK, &SuccessResponseSpec{
		Code:    Success,
		Message: "success",
		Data:    []string{},
	}
}

func SuccessResponseWithData(data interface{}) (int, *SuccessResponseSpec) {
	return http.StatusOK, &SuccessResponseSpec{
		Code:    Success,
		Message: "success",
		Data:    data,
	}
}
