package grpc

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/adapter/models"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/channels"
	"log"
	"reflect"
)

var grpcChannel = channels.GetGrpcChannel()

func ProcessMessage() {
	message := <- grpcChannel
	switch reflect.TypeOf(message).String() {
	case reflect.TypeOf(&models.CreateUserMessage{}).String():
		log.Printf("Create User!!")
	}
}
