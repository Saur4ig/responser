package responser

import (
	"encoding/json"
)

// response body
type ByteResponse []byte

// BlankResponse first of all we need to decide, which struct to use, success true/false will help us
type BlankResponse struct {
	Success bool `json:"success"`
	Data    interface{}
	Error   interface{}
}

// for all success responses
type success struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// for responses with error
type withError struct {
	Success bool  `json:"success"`
	Error   Error `json:"error"`
}

// error response
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Info    string `json:"info"`
}

func (we withError) String() string {
	data, err := json.Marshal(we)
	if err != nil {
		return ""
	}
	return string(data)
}
