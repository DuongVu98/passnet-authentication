package app

import (
	"fmt"
	myproto "github.com/DuongVu98/passnet-authentication/src/main/gen/src/main/proto"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/config"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/repository"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/repository/impl"
	"github.com/Kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func NewAppConfig(sagaMessageClient myproto.MessageServiceClient) *config.AppConfig {
	return &config.AppConfig{MessageServiceClient: sagaMessageClient}
}

func SetMgmClient() {
	dbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	dbName := os.Getenv("DB_NAME")
	err := mgm.SetDefaultConfig(nil, dbName, options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatalf("error when connecting to database --> %s", err.Error())
	}
}

func GetUserMgmRepository() repository.IUserRepository {
	return impl.NewUserMgmRepository()
}

var appConfigInstance = NewAppConfig(GetSagaMessageGrpcClient())

func GetAppConfigInstance() *config.AppConfig {
	return appConfigInstance
}
