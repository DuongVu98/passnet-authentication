package grpc

import (
	"context"
	"github.com/DuongVu98/passnet-authentication/src/main/gen/src/main/proto"
	"github.com/DuongVu98/passnet-authentication/src/main/go/usecases/factory"
	"github.com/DuongVu98/passnet-authentication/src/main/go/usecases/service"
	"log"
)

type CompensatingController struct {
	proto.CompensatingExecutorServer
}

func (server *CompensatingController) Rollback(ctx context.Context, eventId *proto.EventIdProtobuf) (*proto.ServiceResponseProtobuf, error) {
	log.Printf("CompensatingController: rollback command for event: [%v]", eventId.GetEventId())

	var compensatingBackupService = service.GetCompensatingBackupService()
	var compensatingCommand = compensatingBackupService.Find(eventId.GetEventId())

	if compensatingCommand == nil {
		log.Printf("Command not found in the backup for event [%s]", eventId.GetEventId())
	}

	compensatingBackupService.Remove(eventId.GetEventId())

	var compensatingFactory = factory.GetCompensatingExecutorFactory()
	var executor, factoryError = compensatingFactory.Produce(compensatingCommand)
	if factoryError != nil {
		log.Panic(factoryError)
	}

	executor.Rollback(ctx, compensatingCommand)

	return &proto.ServiceResponseProtobuf{
		Message: "SUCCESS",
	}, nil
}
