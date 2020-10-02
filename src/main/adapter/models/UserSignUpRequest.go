package models

type (
	UserSignUpRequest struct {
		Email    string
		Password string
	}

	UserSignUpRequestBuilder interface {
		SetEmail(email string) UserSignUpRequestBuilder
		SetPassword(password string) UserSignUpRequestBuilder
		Build() *UserSignUpRequest
	}
)

func (u UserSignUpRequest) SetEmail(email string) UserSignUpRequestBuilder {
	u.Email = email
	return &u
}

func (u UserSignUpRequest) SetPassword(password string) UserSignUpRequestBuilder {
	u.Password = password
	return &u
}

func (u UserSignUpRequest) Build() *UserSignUpRequest {
	return &u
}

func NewUserSignUpRequestBuilder() UserSignUpRequestBuilder {
	return &UserSignUpRequest{}
}



