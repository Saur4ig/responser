package responser

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

// get response only with success param and all other
func (bytes ByteResponse) GetBlankResponse() (bool, *BlankResponse, error) {
	decider := BlankResponse{}
	if err := json.Unmarshal(bytes, &decider); err != nil {
		return false, nil, errors.Wrap(err, fmt.Sprintf("failed Unmarshal data in blank response struct, RAW - %s", bytes))
	}
	if !decider.Success {
		if decider.Error == nil {
			decider.Success = true
		}
	}

	return decider.Success, &decider, nil
}

// get error struct from response
func (b *BlankResponse) GetError() (*Error, error) {
	e, ok := b.Error.(Error)
	if b.Error == nil || !ok {
		return nil, fmt.Errorf("error data in response not found")
	}
	return &e, nil
}

// get data as interface and check supposed type(not pointer)
func (b *BlankResponse) GetData(requestedDataType interface{}) (interface{}, bool) {
	if requestedDataType == nil || b.Data == nil || !b.Success {
		return nil, false
	}
	dataType := reflect.TypeOf(requestedDataType)
	typeFromResponse := reflect.TypeOf(b.Data)

	if !reflect.DeepEqual(typeFromResponse, dataType) {
		return nil, false
	}

	return b.Data, true
}
