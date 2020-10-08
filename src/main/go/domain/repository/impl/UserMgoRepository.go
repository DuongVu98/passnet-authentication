package impl

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/entity"
)

type UserMgoRepository struct {

}

func NewUserMgoRepository() *UserMgoRepository {
	return &UserMgoRepository{}
}

func (u UserMgoRepository) FindUserByEmailOrDisplayName(emailOrDisplayName string) (*entity.UserEntity, error) {
	panic("implement me")
}

func (u UserMgoRepository) InsertUser(userEntity *entity.UserEntity) (*entity.UserEntity, error) {
	panic("implement me")
}

func (u UserMgoRepository) UpdateUser(userEntity *entity.UserEntity) (*entity.UserEntity, error) {
	panic("implement me")
}



