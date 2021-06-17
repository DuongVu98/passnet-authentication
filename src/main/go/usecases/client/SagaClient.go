package client

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/event"
	"log"
)


type SagaClient struct {

}

func (client SagaClient) Send(event event.UserRegisteredEvent) {
	log.Printf("send event %v", event)
}
