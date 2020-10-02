package repository

import (
	"github.com/DuongVu98/passnet-authentication/src/main/domain/entity"
)

type IUserRepository interface {
	FindUserByEmailOrDisplayName(emailOrDisplayName string) (*entity.UserEntity, error)
	InsertUser(userEntity *entity.UserEntity) (*entity.UserEntity, error)
	UpdateUser(userEntity *entity.UserEntity) (*entity.UserEntity, error)
	LogTest(message string)
}
