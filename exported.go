package responser

import (
	"net/http"
)

var res = New()

// enables logging for every outgoing response
func EnableLogInAllResponses() {
	res.EnableLogInAllResponses()
}

// sends json body with success struct
func SendSuccessResponse(w http.ResponseWriter, body interface{}) error {
	return res.SendSuccessResponse(w, body)
}

// sends only status with empty body
func SendEmptyResponse(w http.ResponseWriter, httpStatus int) {
	res.SendEmptyResponse(w, httpStatus)
}

// sends json body with error struct and status 200
func SendErrorResponse(w http.ResponseWriter, httpStatus int, message string) error {
	return res.SendErrorResponse(w, httpStatus, message)
}
