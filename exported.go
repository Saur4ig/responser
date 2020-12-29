package responser

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
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

// read all from body as BlankResponse
// bool - is request was successful
// BlankResponse - struct with data and error at the same time
// errors possible
func ReadAsBlank(body io.Reader) (bool, *BlankResponse, error) {
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return false, nil, err
	}

	decider := BlankResponse{}
	if err := json.Unmarshal(bytes, &decider); err != nil {
		return false, nil, errors.Wrap(err, fmt.Sprintf("failed Unmarshal data in blank response struct, RAW - %s", bytes))
	}

	if !decider.Success || decider.Error == nil {
		decider.Success = true
	}
	return decider.Success, &decider, nil
}
