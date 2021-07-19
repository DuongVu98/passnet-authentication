package rest

import (
	"context"
	"fmt"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/command"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/form"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/response"
	"github.com/DuongVu98/passnet-authentication/src/main/go/usecases/factory"
	"github.com/labstack/echo/v4"
	"net/http"
)

var commandExecutorFactory = factory.GetCommandExecutorFactory()

func Register(c echo.Context) (err error) {
	var requestContext, cancel = context.WithCancel(context.Background())
	defer cancel()

	m := echo.Map{}
	err = c.Bind(&m)
	if err != nil {
		return err
	}

	registerForm := form.RegisterForm{
		Email:          fmt.Sprintf("%v", m["email"]),
		Username:       fmt.Sprintf("%v", m["username"]),
		Password:       fmt.Sprintf("%v", m["password"]),
		FirstName:      fmt.Sprintf("%v", m["firstName"]),
		LastName:       fmt.Sprintf("%v", m["lastName"]),
		ProfileRole:    fmt.Sprintf("%v", m["role"]),
		OrganizationId: fmt.Sprintf("%v", m["organizationId"]),
		DepartmentId:   fmt.Sprintf("%v", m["departmentId"]),
		CardId:         fmt.Sprintf("%v", m["cardId"]),
	}

	cm := command.RegisterCommand{
		Username:       registerForm.Username,
		Email:          registerForm.Email,
		Password:       registerForm.Password,
		FirstName:      registerForm.FirstName,
		LastName:       registerForm.LastName,
		OrganizationId: registerForm.OrganizationId,
		DepartmentId:   registerForm.DepartmentId,
		CardId:         registerForm.CardId,
		ProfileRole:    registerForm.ProfileRole,
	}

	cme := commandExecutorFactory.Produce(cm)
	_, executorError := cme.Execute(requestContext, cm)

	if executorError != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: executorError.Error()})
	} else {
		return c.JSON(http.StatusOK, "OK")
	}
}
