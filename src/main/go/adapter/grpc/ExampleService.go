package grpc

import (
	"context"
	myproto "github.com/DuongVu98/passnet-authentication/src/main/gen/src/main/proto"
	"log"
)

type ExampleService struct {
}

func (s *ExampleService) SayHello(ctx *context.Context, message *myproto.MyMessage) (*myproto.MyMessage, error) {
	log.Printf("Receive message from client --> %s", message.Body)
	return &myproto.MyMessage{Body: "Hello client! This is server"}, nil
}
