package model

import (
	"fmt"
)

//Error Ошибка
//swagger:model Error
type Error struct {
	Message string `json:"error"`
}

//NewError создание "Ошибки"
func NewError(msg string) *Error {
	fmt.Println(msg)
	return &Error{Message: msg}
}

//Error реализуем интерфес error
func (e *Error) Error() string {
	return fmt.Sprintf("%s", e.Message)
}
