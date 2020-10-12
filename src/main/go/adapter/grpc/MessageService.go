package grpc

import (
	"context"
	"github.com/DuongVu98/passnet-authentication/src/main/gen/src/main/proto"
	"github.com/DuongVu98/passnet-authentication/src/main/go/adapter/models"
	"github.com/DuongVu98/passnet-authentication/src/main/go/config/app"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/channels"
	"log"
	"reflect"
)

var grpcChannel = channels.GetGrpcChannel()

var appConfig = app.GetAppConfigInstance()
var sagaMessageClient = appConfig.MessageServiceClient

func ProcessMessage() {
	message := <-grpcChannel
	switch reflect.TypeOf(message).String() {
	case reflect.TypeOf(&models.CreateUserMessage{}).String():
		log.Printf("Create User!!")
		_, _ = sagaMessageClient.SendCreateUserMessage(context.Background(), &proto.CreateUserMessage{Uid: message.(*models.CreateUserMessage).Uid})
	}
}

