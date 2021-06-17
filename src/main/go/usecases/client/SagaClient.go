package client

import (
	"context"
	gen "github.com/DuongVu98/passnet-authentication/src/main/gen/src/main/proto"
	"github.com/DuongVu98/passnet-authentication/src/main/go/config/app"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/event"
	"log"
)

var grpcClient = app.GetSagaEventProducerClient()

type SagaClient struct {

}

func GetSagaClient() SagaClient {
	return SagaClient{}
}

func (client SagaClient) Send(event event.UserRegisteredEvent) {
	log.Printf("send event %v", event)
	var message = gen.UserRegisteredEvent{
		Uid: event.Uid,
		Username: event.Username,
		Email: event.Email,
		FirstName: event.FirstName,
		LastName: event.LastName,
	}
	var response, err = grpcClient.ProduceUserRegisteredEvent(context.Background(), &message)

	if err != nil {
		log.Printf("err during client call grpc %v", err)
	}
	log.Printf("message from grpc server %v", response)
}
