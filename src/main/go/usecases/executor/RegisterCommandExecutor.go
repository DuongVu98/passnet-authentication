package executor

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/config/app"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/command"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"log"
)

type RegisterCommandExecutor struct {
}

func (e RegisterCommandExecutor) Execute(c command.RegisterCommand)  {
	oktaClient := app.OktaClient()

	uc := &okta.UserCredentials{
		Password: &okta.PasswordCredential{
			Value: c.Password,
		},
	}
	profile := okta.UserProfile{}
	profile["email"] = c.Email
	profile["login"] = c.Email
	profile["firstName"] = c.FirstName
	profile["lastName"] = c.LastName
	request := okta.CreateUserRequest{Credentials: uc, GroupIds: []string{}, Profile: &profile, Type: &okta.UserType{}}

	user, resp, err := oktaClient.User.CreateUser(app.OktaContext(), request, nil)

	if err != nil {
		log.Printf("Failed: %s", err)
	}

	log.Printf("repsonse created user: %s", resp.NextPage)
	log.Printf("user created: %v", user)
}
