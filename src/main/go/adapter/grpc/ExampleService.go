package grpc

import (
	"context"
	"github.com/DuongVu98/passnet-authentication/src/main/gen/src/main/proto"
	"log"
)

type ExampleService struct {
}

func (s *ExampleService) SayHello(ctx *context.Context, message *proto.MyMessage) (*proto.MyMessage, error) {
	log.Printf("Receive message from client --> %s", message.Body)
	return &proto.MyMessage{Body: "Hello client! This is server"}, nil
}
