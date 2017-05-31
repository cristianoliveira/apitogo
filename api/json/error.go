package json

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	Code   int      `json:"code"`
	Errors []string `json:"errors"`
}

func (e *Error) AsBytes() []byte {
	data, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func NewError(code int, err error) *Error {
	return &Error{Errors: []string{err.Error()}, Code: code}
}
