package impl

import (
	"errors"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/channels"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/config"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/entity"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/repository"
	"gopkg.in/mgo.v2/bson"
	"log"
)

var appConfigChannel = channels.GetAppConfigChannel()

type UserBongoRepository struct {
}

func (u *UserBongoRepository) FindUserByUid(uid string) (*entity.UserEntity, error) {
	panic("implement me")
}

func NewUserBongoRepository() repository.IUserRepository {
	return &UserBongoRepository{}
}
func (u *UserBongoRepository) FindUserByEmailOrDisplayName(emailOrDisplayName string) (*entity.UserEntity, error) {

	appConfig := <-appConfigChannel
	bongoClient := appConfig.(config.AppConfig).BongoClient

	userEntity := &entity.UserEntity{}
	err1 := bongoClient.Collection("user").FindOne(bson.M{"email": emailOrDisplayName}, userEntity)
	err2 := bongoClient.Collection("user").FindOne(bson.M{"display_name": emailOrDisplayName}, userEntity)
	if err1 != nil && err2 != nil {
		log.Printf("email and display_name not found")
		err := errors.New("email and display_name not found")
		return nil, err
	}
	return userEntity, nil
}

func (u *UserBongoRepository) InsertUser(userEntity *entity.UserEntity) (*entity.UserEntity, error) {
	//log.Printf("hello from bongo")
	//err := bongoClient.Collection("user").Save(bongo.Document(userEntity))
	//
	//log.Printf("InserUser")
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//signUpUserEntity := &entity.UserEntity{}
	//err1 := bongoClient.Collection("user").FindOne(bson.M{"email": userEntity.Email}, signUpUserEntity)
	//err2 := bongoClient.Collection("user").FindOne(bson.M{"display_name": userEntity.DisplayName}, signUpUserEntity)
	//if err1 == nil || err2 == nil {
	//	return signUpUserEntity, nil
	//}
	//return nil, errors.New("some error from get user after sign up")
	return nil, nil
}

func (u *UserBongoRepository) UpdateUser(userEntity *entity.UserEntity) (*entity.UserEntity, error) {
	panic("implement me")
}

func (u *UserBongoRepository) LogTest(message string) {
	panic(message)
}
