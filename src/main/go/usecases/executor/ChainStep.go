package executor

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/aggregate"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/command"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/event"
	client2 "github.com/DuongVu98/passnet-authentication/src/main/go/usecases/client"
	"reflect"
)

func (step PublishEventStep) Execute(c command.BaseCommand) (aggregate.User, error) {
	switch reflect.TypeOf(c).String() {
	case reflect.TypeOf(command.RegisterCommand{}).String():
		var user, err = step.Executor.Execute(c)
		if err != nil {
			return step.Executor.Execute(c)
		} else {
			var eventToSend = event.UserRegisteredEvent{
				Uid: user.Uid.Value,
				Username: user.Username.Value,
				Email: user.Email.Value,
				FirstName: user.Profile.FirstName,
				LastName: user.Profile.LastName,
			}

			var client = client2.GetSagaClient()
			client.Send(eventToSend)
			return user, nil
		}
	default:
		return step.Executor.Execute(c)
	}
}
