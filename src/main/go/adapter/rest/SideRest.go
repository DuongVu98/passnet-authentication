package rest

import (
	"fmt"
	"github.com/DuongVu98/passnet-authentication/src/main/go/config/app"
	"github.com/labstack/echo/v4"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
	"log"
	"net/http"
)

func DeleteUser(c echo.Context) (err error) {
	email := c.QueryParam("email")
	oktaClient := app.OktaClient()
	oktaContext := app.OktaContext()

	//fmt.Println(fmt.Sprintf("firstName eq \"%s\"", email))
	filter := query.NewQueryParams(query.WithFilter("status eq \"ACTIVE\""))
	users, _, _ := oktaClient.User.ListUsers(oktaContext, filter)

	var i = 0
	if len(users) > 0 {
		i++
		for _, u := range users {

			var _, err1 = oktaClient.User.DeactivateUser(app.OktaContext(), u.Id, query.NewQueryParams())
			if err1 != nil {
				log.Panic(err1)
			}

			var _, err2 = oktaClient.User.DeactivateOrDeleteUser(app.OktaContext(), u.Id, query.NewQueryParams())
			if err2 != nil {
				log.Panic(err2)
			}
		}
	}
	return c.String(http.StatusOK, fmt.Sprintf("email %s deleted with %v users", email, i))
}