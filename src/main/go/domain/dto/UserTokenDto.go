package dto

type (
	UserTokenDto struct {
		UserDto *UserDto `json:"user_dto"`
		Token   string   `json:"token"`
	}
	UserTokenDtoBuilder interface {
		SetUserDto(userDto *UserDto) UserTokenDtoBuilder
		SetToken(token string) UserTokenDtoBuilder
		Build() *UserTokenDto
	}
)

func (u UserTokenDto) SetUserDto(userDto *UserDto) UserTokenDtoBuilder {
	u.UserDto = userDto
	return u
}

func (u UserTokenDto) SetToken(token string) UserTokenDtoBuilder {
	u.Token = token
	return u
}

func (u UserTokenDto) Build() *UserTokenDto {
	return &u
}

func NewUserTokenDtoBuilder() UserTokenDtoBuilder {
	return &UserTokenDto{}
}



