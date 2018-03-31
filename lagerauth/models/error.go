package models

type Error struct {
	IsError bool   `json:"error"`
	Message string `json:"message"`
}

func NewError(err error) Error {
	return Error{IsError: true, Message: err.Error()}
}
