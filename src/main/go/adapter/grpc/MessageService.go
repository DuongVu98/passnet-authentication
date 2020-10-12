package grpc

import (
	"context"
	"github.com/DuongVu98/passnet-authentication/src/main/gen/src/main/proto"
	"github.com/DuongVu98/passnet-authentication/src/main/go/adapter/models"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/channels"
	"google.golang.org/grpc"
	"log"
	"reflect"
)

var grpcChannel = channels.GetGrpcChannel()

func ProcessMessage() {
	message := <-grpcChannel
	switch reflect.TypeOf(message).String() {
	case reflect.TypeOf(&models.CreateUserMessage{}).String():
		log.Printf("Create User!!")
	}
}

type MessageServiceClient struct {
}

func (m MessageServiceClient) SendCreateUserMessage(ctx context.Context, in *proto.CreateUserMessage, opts ...grpc.CallOption) (*proto.Response, error) {
	panic("implement me")
}

