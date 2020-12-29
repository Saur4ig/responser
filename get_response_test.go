package responser

import (
	"reflect"
	"testing"
)

func TestBlankResponse_GetData(t *testing.T) {
	type fields struct {
		Success bool
		Data    interface{}
		Error   interface{}
	}
	type test1 struct {
		One int
		Two string
	}
	type test2 struct {
		Test float32
		Some bool
	}
	tests := []struct {
		name   string
		fields fields
		data   interface{}
		wantOK bool
	}{
		{
			name: "all ok example",
			fields: fields{
				Success: true,
				Data: test1{
					One: 1,
					Two: "ss",
				},
				Error: nil,
			},
			data:   test1{},
			wantOK: true,
		},
		{
			name: "example with wrong type",
			fields: fields{
				Success: true,
				Data: test1{
					One: 2,
					Two: "pp",
				},
				Error: nil,
			},
			data:   test2{},
			wantOK: false,
		},
		{
			name: "example with error, without data",
			fields: fields{
				Success: true,
				Data:    nil,
				Error: Error{
					Code:    404,
					Message: "not found at all",
					Info:    "something",
				},
			},
			data:   test1{},
			wantOK: false,
		},
		{
			name: "without required type",
			fields: fields{
				Success: true,
				Data: test1{
					One: 1,
					Two: "22",
				},
				Error: nil,
			},
			data:   nil,
			wantOK: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BlankResponse{
				Success: tt.fields.Success,
				Data:    tt.fields.Data,
				Error:   tt.fields.Error,
			}
			_, ok := b.GetData(tt.data)
			if ok != tt.wantOK {
				t.Errorf("GetData() failed, something went wrong")
			}
		})
	}
}

func TestBlankResponse_GetError(t *testing.T) {
	type fields struct {
		Success bool
		Data    interface{}
		Error   interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Error
		wantErr bool
	}{
		{
			name: "example with normal error",
			fields: fields{
				Success: false,
				Data:    nil,
				Error: Error{
					Code:    500,
					Message: "all bad",
					Info:    "",
				},
			},
			want: &Error{
				Code:    500,
				Message: "all bad",
				Info:    "",
			},
			wantErr: false,
		},
		{
			name: "without error",
			fields: fields{
				Success: false,
				Data:    nil,
				Error:   nil,
			},
			wantErr: true,
		},
		{
			name: "successful response",
			fields: fields{
				Success: true,
				Data: struct {
					one string
				}{one: "test1"},
				Error: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BlankResponse{
				Success: tt.fields.Success,
				Data:    tt.fields.Data,
				Error:   tt.fields.Error,
			}
			respErr, err := b.GetError()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(respErr, tt.want) {
				t.Errorf("GetError() got = %v, want %v", respErr, tt.want)
			}
		})
	}
}
