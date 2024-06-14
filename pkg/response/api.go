package response

import "net/http"

type API struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Error   error       `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseDBError(err error, message string) (API, error) {
	if message == "" {
		message = err.Error()
	}

	return API{
		Status:  http.StatusInternalServerError,
		Error:   err,
		Message: message,
	}, err
}

func Forbidden(message string) (API, error) {
	return API{
		Status:  http.StatusForbidden,
		Message: message,
	}, nil
}

func BadRequest(err error, message string) (API, error) {
	if message == "" {
		message = err.Error()
	}

	return API{
		Status:  http.StatusBadRequest,
		Message: message,
		Error:   err,
	}, err
}

func NotFound(message string) (API, error) {
	return API{
		Status:  http.StatusNotFound,
		Message: message,
	}, nil
}
