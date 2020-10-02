package impl

import (
	"errors"
	"github.com/DuongVu98/passnet-authentication/src/main/domain/entity"
	"github.com/Kamva/mgm/v3"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type UserMgmRepository struct {
}

func NewUserMgmRepository() *UserMgmRepository {
	return &UserMgmRepository{}
}

func (u *UserMgmRepository) FindUserByEmailOrDisplayName(emailOrDisplayName string) (*entity.UserEntity, error) {
	userCollection := mgm.Coll(&entity.UserEntity{})
	userByEmail := &entity.UserEntity{}
	userByDisplayName := &entity.UserEntity{}
	err1 := userCollection.First(bson.M{"email": emailOrDisplayName}, userByEmail)
	err2 := userCollection.First(bson.M{"display_name": emailOrDisplayName}, userByDisplayName)

	if err1 == nil {
		return userByEmail, nil
	} else if err2 == nil {
		return userByDisplayName, nil
	} else {
		return nil, errors.New("not found")
	}
}

func (u *UserMgmRepository) InsertUser(userEntity *entity.UserEntity) (*entity.UserEntity, error) {
	log.Println(" before creat collection")
	userCollection := mgm.Coll(userEntity)
	log.Println(" after creat collection")

	err := userCollection.Create(userEntity)
	if err != nil {
		return nil, err
	}
	insertedUser := &entity.UserEntity{}
	err = userCollection.First(bson.M{"email": userEntity.Email}, insertedUser)
	if err != nil {
		log.Printf("error when insert user: user after insert not found --> %s", err.Error())
		return nil, err
	}
	log.Printf("get user after insert")
	return insertedUser, nil
}

func (u *UserMgmRepository) UpdateUser(userEntity *entity.UserEntity) (*entity.UserEntity, error) {
	userCollection := mgm.Coll(&entity.UserEntity{})
	err := userCollection.Update(userEntity)
	if err != nil {
		return nil, err
	}
	insertedUser := &entity.UserEntity{}
	_ = userCollection.First(bson.M{"email": userEntity.Email}, insertedUser)

	return insertedUser, nil
}

func (u *UserMgmRepository) LogTest(message string) {
	log.Printf(message)
}
