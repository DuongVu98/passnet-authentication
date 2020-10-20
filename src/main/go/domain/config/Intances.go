package config

import (
	myproto "github.com/DuongVu98/passnet-authentication/src/main/gen/src/main/proto"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/repository"
)

type (
	AppConfig struct {
		MessageServiceClient myproto.MessageServiceClient
	}
	BeanConfig struct {
		UserRepository repository.IUserRepository
	}
	SingletonFactory struct {
		factory map[string]interface{}
	}
)

func (s *SingletonFactory) Set(key string, value interface{}) {
	s.factory[key] = value
}
func (s *SingletonFactory) Get (key string) interface{} {
	return s.factory[key]
}

/*
Create instance
 */
var singletonFactory = &SingletonFactory{
	factory: make(map[string]interface{}),
}
func GetSingletonFactory() *SingletonFactory {
	return singletonFactory
}
