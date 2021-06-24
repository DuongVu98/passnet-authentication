package chain

import (
	"context"
	"github.com/DuongVu98/passnet-authentication/src/main/go/config/app"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/aggregate"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/command"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/event"
	client2 "github.com/DuongVu98/passnet-authentication/src/main/go/usecases/client"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
	"log"
	"reflect"
)

func (step PublishEventStep) Execute(requestContext context.Context,c command.BaseCommand) (aggregate.User, error) {
	switch reflect.TypeOf(c).String() {
	case reflect.TypeOf(command.RegisterCommand{}).String():
		var user, err = step.Executor.Execute(requestContext, c)
		if err != nil {
			return step.Executor.Execute(requestContext, c)
		} else {
			var eventToSend = event.UserRegisteredEvent{
				Uid:       user.Uid.Value,
				Username:  user.Username.Value,
				Email:     user.Email.Value,
				FirstName: user.Profile.FirstName,
				LastName:  user.Profile.LastName,
			}

			var client = client2.GetSagaClient()
			var eventId = requestContext.Value("eventId").(string)
			var grpcClientError = client.Send(eventToSend, eventId)

			if grpcClientError != nil {
				rollbackCreateUserLocalTransaction(user.Uid.Value)
				return user, grpcClientError
			} else {
				return user, nil
			}
		}
	default:
		return step.Executor.Execute(requestContext, c)
	}
}

func rollbackCreateUserLocalTransaction(userId string)  {
	var oktaClient = app.OktaClient()
	var _, _ = oktaClient.User.DeactivateUser(context.Background(), userId, query.NewQueryParams())
	var resp, err = oktaClient.User.DeactivateOrDeleteUser(context.Background(), userId, query.NewQueryParams())
	if err != nil {
		log.Panicln(err)
	} else {
		log.Println("Grpc client error, about to delete created user...")
		log.Printf("Transactional success: %v", resp.Status)
	}
}
