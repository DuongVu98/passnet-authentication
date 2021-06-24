package chain

import (
	"context"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/aggregate"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/command"
	"github.com/DuongVu98/passnet-authentication/src/main/go/usecases/service"
	"log"
	"reflect"
)

func (step BackupCompensatingStep) Execute(requestContext context.Context, c command.BaseCommand) (aggregate.User, error) {

	var ev = requestContext.Value("eventId")
	log.Printf("BackupCompensatingStep: get eventId from context: %s", ev.(string))

	switch reflect.TypeOf(c).String() {
	case reflect.TypeOf(command.RegisterCommand{}).String():
		var result, err = step.Executor.Execute(requestContext, c)

		if err != nil {
			log.Printf("BackupCompensatingStep: Failed - %s", err)
			return aggregate.User{}, err
		}

		var compensatingCommand = command.RegisterCommandCompensating{
			AggregatId: result.Uid.Value,
		}

		var compensatingBackupService = service.GetCompensatingBackupService()
		var eventId = requestContext.Value("eventId")
		log.Printf("BackupCompensatingStep: get eventId from context: %s", eventId)

		compensatingBackupService.Store(compensatingCommand, eventId.(string))
		return result, nil
	default:
		return step.Executor.Execute(requestContext, c)
	}
}