package chain

import (
	"context"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/aggregate"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/command"
	"github.com/DuongVu98/passnet-authentication/src/main/go/usecases/service"
	"reflect"
)

func (step BackupCompensatingStep) Execute(requestContext context.Context, c command.BaseCommand) (aggregate.User, error) {

	switch reflect.TypeOf(c).String() {
	case reflect.TypeOf(command.RegisterCommand{}).String():
		var result, _ = step.Execute(requestContext, c)

		var compensatingCommand = command.RegisterCommandCompensating{
			AggregatId: result.Uid.Value,
		}

		var compensatingBackupService = service.GetCompensatingBackupService()
		var eventId = requestContext.Value("eventId").(string)

		compensatingBackupService.Store(compensatingCommand, eventId)
		return result, nil
	default:
		return step.Execute(requestContext, c)
	}
}