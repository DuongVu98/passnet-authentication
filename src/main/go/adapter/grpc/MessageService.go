package grpc

import (
	"context"
	"github.com/DuongVu98/passnet-authentication/src/main/gen/src/main/proto"
	"github.com/DuongVu98/passnet-authentication/src/main/go/adapter/models"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/channels"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/config"
	"log"
	"reflect"
)


func ProcessMessage() {
	var grpcChannel = channels.GetGrpcChannel()
	var singletonFactory = config.GetSingletonFactory()
	var appConfig = singletonFactory.Get("appConfig").(*config.AppConfig)
	var sagaMessageClient = appConfig.MessageServiceClient

	message := <-grpcChannel
	switch reflect.TypeOf(message).String() {
	case reflect.TypeOf(&models.CreateUserMessage{}).String():
		log.Printf("Create User!!")
		messageToSend := &proto.CreateUserMessage{Uid: message.(*models.CreateUserMessage).Uid}
		response, err := sagaMessageClient.SendCreateUserMessage(context.Background(), messageToSend)
		if err != nil {
			log.Printf("err during client call grpc %v", err)
		}
		log.Printf("message from grpc server %s", response.Message)
	}
}

