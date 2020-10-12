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
)
