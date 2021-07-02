package grpc

import (
	"context"
	"github.com/DuongVu98/passnet-authentication/src/main/gen/src/main/proto"
	"log"
)

type CompensatingController struct {
	proto.CompensatingExecutorServer
}

func (server *CompensatingController) Rollback(ctx context.Context, eventId *proto.EventIdProtobuf) (*proto.ServiceResponseProtobuf, error) {
	log.Printf("CompensatingController: rollback command for event: [%v]", eventId.GetEventId())
	return &proto.ServiceResponseProtobuf{
		Message: "SUCCESS",
	}, nil
}
