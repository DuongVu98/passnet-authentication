package config

import (
	firebase "firebase.google.com/go/v4"
	myproto "github.com/DuongVu98/passnet-authentication/src/main/gen/src/main/proto"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/repository"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2"
)

type (
	AppConfig struct {
		FirebaseApp          firebase.App
		BongoClient          *bongo.Connection
		MgoClient            *mgo.Session
		MessageServiceClient myproto.MessageServiceClient
	}
	BeanConfig struct {
		UserRepository repository.IUserRepository
	}
)
