package model

import (
	"github.com/go-openapi/strfmt"
)

//TokenInfo получение информации от oauth
type TokenInfo struct {
	ClientID string
	UserID   string
	Scope    string
	Token    string
}

//GetClientID получить идентификатор приложения
func (ti *TokenInfo) GetClientID() (clientID strfmt.UUID4, err error) {
	clientID = strfmt.UUID4(ti.ClientID)
	return
}

//GetUserID получить идентификатор пользователя
func (ti *TokenInfo) GetUserID() (userID strfmt.UUID4, err error) {
	userID = strfmt.UUID4(ti.UserID)
	return
}

//UserIsNull проверить есть ли пользователь
func (ti *TokenInfo) UserIsNull() bool {
	if ti.UserID == "" {
		return true
	}

	return false
}
