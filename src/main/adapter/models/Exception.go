package models

type Exception struct {
	Message string `json:"message"`
}

func NewException(message string) *Exception {
	return &Exception{Message: message}
}
