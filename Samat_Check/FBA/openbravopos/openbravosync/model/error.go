package model

import (
	"fmt"
)

type Error struct {
	Error string `json:"error"`
}

func NewError(msg string) *Error {
	fmt.Println(msg)
	return &Error{Error: msg}
}
