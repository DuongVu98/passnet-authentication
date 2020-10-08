package models

type (
	UserLoginRequest struct {
		EmailOrDisplayName string `json:"emailOrDisplayName"`
		Password           string `json:"password"`
	}
	UserLoginRequestBuilder interface {
		SetEmailOrDisplayName(emailOrDisplayName string) UserLoginRequestBuilder
		SetPassword(password string) UserLoginRequestBuilder
		Build() *UserLoginRequest
	}
)

func NewUserLoginRequest(emailOrDisplayName string, password string) *UserLoginRequest {
	return &UserLoginRequest{EmailOrDisplayName: emailOrDisplayName, Password: password}
}

func (u UserLoginRequest) SetEmailOrDisplayName(emailOrDisplayName string) UserLoginRequestBuilder {
	u.EmailOrDisplayName = emailOrDisplayName
	return &u
}

func (u UserLoginRequest) SetPassword(password string) UserLoginRequestBuilder {
	u.Password = password
	return &u
}

func (u UserLoginRequest) Build() *UserLoginRequest {
	return &u
}

func NewUserLoginRequestBuilder() UserLoginRequestBuilder {
	return &UserLoginRequest{}
}
