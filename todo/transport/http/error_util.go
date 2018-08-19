package http

import (
	"golang-tdd-rest-api/todo"
	"net/http"
)

type ResponseError struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
}

func getErrorResponse(err error) (httpStatus int, res ResponseError) {
	httpStatus = http.StatusInternalServerError
	res = ResponseError{}

	switch err {
	case todo.NOT_FOUND_ERROR:
		httpStatus = http.StatusBadRequest
		res.Error.Code = "E001"
		res.Error.Message = "No todo list"
	case todo.WRONG_PARAMS:
		httpStatus = http.StatusBadRequest
		res.Error.Code = "E002"
		res.Error.Message = "Wrong params"
	default:
		httpStatus = http.StatusInternalServerError
		res.Error.Code = "E999"
		res.Error.Message = "Internal Server Error"
	}

	return
}
