package entity

import (
	"github.com/Kamva/mgm/v3"
)

type UserEntity struct {
	//bongo.DocumentBase `bson:",inline"`
	mgm.DefaultModel `bson:",inline"`
	Uid              string             `json:"uid"`
	Email            string             `json:"email"`
	DisplayName      string             `json:"display_name"`
	AvatarUrl        string             `json:"avatar_url"`
	Auth             UserAuthentication `json:"auth"`
}
type UserAuthentication struct {
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

func NewUserAuthentication(password string, salt string) *UserAuthentication {
	return &UserAuthentication{Password: password, Salt: salt}
}

type UserEntityBuilder interface {
	SetUid(uid string) UserEntityBuilder
	SetEmail(email string) UserEntityBuilder
	SetDisplayName(displayName string) UserEntityBuilder
	SetAvatarUrl(avatarUrl string) UserEntityBuilder
	SetAuthentication(auth *UserAuthentication) UserEntityBuilder
	Build() *UserEntity
}

func (u UserEntity) SetUid(uid string) UserEntityBuilder {
	u.Uid = uid
	return u
}

func (u UserEntity) SetEmail(email string) UserEntityBuilder {
	u.Email = email
	return u
}

func (u UserEntity) SetDisplayName(displayName string) UserEntityBuilder {
	u.DisplayName = displayName
	return u
}

func (u UserEntity) SetAvatarUrl(avatarUrl string) UserEntityBuilder {
	u.AvatarUrl = avatarUrl
	return u
}
func (u UserEntity) SetAuthentication(auth *UserAuthentication) UserEntityBuilder {
	u.Auth = *auth
	return u
}

func (u UserEntity) Build() *UserEntity {
	return &u
}

func NewUserEntityBuilder() UserEntityBuilder {
	return &UserEntity{}
}
