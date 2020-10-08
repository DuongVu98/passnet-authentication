package dto

type UserDto struct {
	Uid         string `json:"uid"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
}
func NewUserDto(uid string, displayName string, email string) *UserDto {
	return &UserDto{Uid: uid, DisplayName: displayName, Email: email}
}

type UserDtoBuilder interface {
	SetUid(uid string) UserDtoBuilder
	SetDisplayName(displayName string) UserDtoBuilder
	SetEmail(email string) UserDtoBuilder
	Build() *UserDto
}
func NewUserDtoBuilder() UserDtoBuilder {
	return &UserDto{}
}

func (u UserDto) SetUid(uid string) UserDtoBuilder {
	u.Uid = uid
	return u
}

func (u UserDto) SetDisplayName(displayName string) UserDtoBuilder {
	u.DisplayName = displayName
	return u
}

func (u UserDto) SetEmail(email string) UserDtoBuilder {
	u.Email = email
	return u
}

func (u UserDto) Build() *UserDto {
	return &u
}
