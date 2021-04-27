package model

import "net/http"

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResultData(data interface{}) *Result {
	return &Result{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	}
}
func NewSuccessResult() *Result {
	return &Result{
		Code:    http.StatusOK,
		Message: "success",
	}
}
func NewFailedCodeResult(code int, err error) *Result {
	return &Result{
		Code:    code,
		Message: err.Error(),
	}
}
func NewFailedResult(err error) *Result {
	return &Result{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	}
}
