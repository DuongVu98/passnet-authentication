package executor

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/config/app"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/aggregate"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/command"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/exception"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"log"
	"reflect"
)

type RegisterCommandExecutor struct {
	CommandExecutor
}

func (e RegisterCommandExecutor) Execute(c command.BaseCommand) (aggregate.User, error) {
	switch reflect.TypeOf(c).String() {
	case reflect.TypeOf(command.RegisterCommand{}).String():
		oktaClient := app.OktaClient()
		uc := &okta.UserCredentials{
			Password: &okta.PasswordCredential{
				Value: c.(command.RegisterCommand).Password,
			},
		}
		profile := okta.UserProfile{}
		profile["email"] = c.(command.RegisterCommand).Email
		profile["login"] = c.(command.RegisterCommand).Email
		profile["firstName"] = c.(command.RegisterCommand).FirstName
		profile["lastName"] = c.(command.RegisterCommand).LastName
		request := okta.CreateUserRequest{Credentials: uc, GroupIds: []string{}, Profile: &profile, Type: &okta.UserType{}}

		user, resp, err := oktaClient.User.CreateUser(app.OktaContext(), request, nil)

		if err != nil {
			log.Printf("Failed: %s", err)
			return aggregate.User{}, err
		} else {
			log.Printf("repsonse created user: %v", resp.NextPage)
			log.Printf("user created: %v", user.Id)

			return aggregate.User{
				Uid:        aggregate.UserId{Value: user.Id},
				Username:   aggregate.Username{Value: c.(command.RegisterCommand).Username},
				Email:      aggregate.Email{Value: c.(command.RegisterCommand).Email},
				Profile:    aggregate.UserProfile{FirstName: c.(command.RegisterCommand).FirstName, LastName: c.(command.RegisterCommand).LastName},
				Credential: aggregate.UserCredential{Password: c.(command.RegisterCommand).Password},
			}, nil
		}
	default:
		return aggregate.User{}, exception.InvalidCommandException{}
	}
}
