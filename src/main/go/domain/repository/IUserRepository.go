package repository

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/entity"
)

type IUserRepository interface {
	FindUserByUid(uid string) (*entity.UserEntity, error)
	FindUserByEmailOrDisplayName(emailOrDisplayName string) (*entity.UserEntity, error)
	InsertUser(userEntity *entity.UserEntity) (*entity.UserEntity, error)
	UpdateUser(userEntity *entity.UserEntity) (*entity.UserEntity, error)
	LogTest(message string)
}
