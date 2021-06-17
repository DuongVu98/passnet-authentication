package grpc

import (
	"context"
	"github.com/DuongVu98/passnet-authentication/src/main/gen/src/main/proto"
	"github.com/DuongVu98/passnet-authentication/src/main/go/config/app"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/channels"
	models2 "github.com/DuongVu98/passnet-authentication/src/main/go/domain/models"
	"log"
	"reflect"
)


func ProcessMessage() {
	var grpcChannel = channels.GetGrpcChannel()
	var sagaMessageClient = app.GetSagaMessageGrpcClient()

	message := <-grpcChannel
	switch reflect.TypeOf(message).String() {
	case reflect.TypeOf(&models2.CreateUserMessage{}).String():
		log.Printf("Create User!!")
		messageToSend := &proto.CreateUserMessage{Uid: message.(*models2.CreateUserMessage).Uid}
		response, err := sagaMessageClient.SendCreateUserMessage(context.Background(), messageToSend)
		if err != nil {
			log.Printf("err during client call grpc %v", err)
		}
		log.Printf("message from grpc server %s", response.Message)
	}
}

