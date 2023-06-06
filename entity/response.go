package entity

import "net/http"

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Payload    interface{} `json:"payload,omitempty"`
}

type CustomError struct {
	Code       string `json:"code"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

var (
	badRequestError = CustomError{
		Code:       "BAD_REQUEST",
		StatusCode: http.StatusBadRequest,
	}

	repositoryError = CustomError{
		Code: 		"REPOSITORY_ERROR",
		StatusCode: http.StatusInternalServerError,
	}

	generalError = CustomError{
		Code: 		"INTERNAL_SERVER_ERROR",
		StatusCode: http.StatusInternalServerError,
	}

	notFoundError = CustomError{
		Code: 		"NOT_FOUND_ERROR",
		StatusCode: http.StatusNotFound,
	}

	unauthorizedError = CustomError{
		Code: "UNAUTHORIZED",
		StatusCode: http.StatusUnauthorized,
	}
)

func BadRequestError(message string) *CustomError {
	err := badRequestError
	err.Message = message
	
	return &err
}

func RepositoryError(message string) *CustomError {
	err := repositoryError
	err.Message = message

	return &err
}

func GeneralError(message string) *CustomError {
	err := generalError
	err.Message = message

	return &err
}

func NotFoundError(message string) *CustomError {
	err := notFoundError
	err.Message = message

	return &err
}

func UnauthorizedError(message string) *CustomError {
	err := unauthorizedError
	err.Message = message

	return &err
}