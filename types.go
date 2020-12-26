package responser

import (
	"encoding/json"
)

// for all success responses
type success struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}

// for responses with error
type withError struct {
	Status Status `json:"status"`
	Error  err    `json:"error"`
}

// error response
type err struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Info    interface{} `json:"info"`
}

func (we withError) String() string {
	data, err := json.Marshal(we)
	if err != nil {
		return ""
	}
	return string(data)
}
