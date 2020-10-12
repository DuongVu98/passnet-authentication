package gateway

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/channels"
)

type MessageGateway struct {
}

func NewMessageGateway() *MessageGateway {
	return &MessageGateway{}
}

var grpcChannel = channels.GetGrpcChannel()

func (mg *MessageGateway) SendMessage(message interface{}) {
	grpcChannel <- message
}
