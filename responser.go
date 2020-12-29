package responser

import (
	"encoding/json"
	"log"
	"net/http"
)

type Responser struct {
	Log bool
}

func New() *Responser {
	return &Responser{
		Log: false,
	}
}

func (r *Responser) EnableLogInAllResponses() {
	r.Log = true
}

func (r *Responser) SendSuccessResponse(w http.ResponseWriter, body interface{}) error {
	// create response body
	response := success{
		Success: true,
		Data:    body,
	}

	return r.sendResponse(w, http.StatusOK, response)
}

func (r *Responser) SendEmptyResponse(w http.ResponseWriter, httpStatus int) {
	w.WriteHeader(httpStatus)
}

func (r *Responser) SendErrorResponse(w http.ResponseWriter, httpStatus int, message string) error {
	// create response body
	response := withError{
		Success: false,
		Error: Error{
			Code:    httpStatus,
			Message: message,
		},
	}

	return r.sendResponse(w, http.StatusOK, response)
}

// sends json body with status
func (r *Responser) sendResponse(w http.ResponseWriter, httpStatus int, body interface{}) error {
	data, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	if r.Log {
		log.Println("response_log", data)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	_, err = w.Write(data)
	return err
}
