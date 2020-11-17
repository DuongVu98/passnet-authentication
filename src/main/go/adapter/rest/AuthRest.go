package rest

import (
	"fmt"
	"github.com/DuongVu98/passnet-authentication/src/main/go/adapter/models"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/channels"
	"github.com/DuongVu98/passnet-authentication/src/main/go/usecases/interactors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

var beanConfigChannel = channels.GetBeanConfigChannel()

func Login(c echo.Context) (err error) {

	m := echo.Map{}
	err = c.Bind(&m)
	if err != nil {
		return err
	}
	loginRequest := models.NewUserLoginRequestBuilder().
		SetEmailOrDisplayName(fmt.Sprintf("%v", m["emailOrDisplayName"])).
		SetPassword(fmt.Sprintf("%v", m["password"])).
		Build()

	loginInteractor := interactors.NewUserLoginInteractor(loginRequest.EmailOrDisplayName, loginRequest.Password)
	userTokenDto, err := loginInteractor.Execute()
	if err != nil {
		log.Print("err 1")
		return c.JSON(http.StatusOK, models.NewException(err.Error()))
	}

	return c.JSON(http.StatusOK, userTokenDto)
}

func SignUp(c echo.Context) (err error) {
	m := echo.Map{}
	err = c.Bind(&m)
	if err != nil {
		return err
	}
	signUpRequest := models.NewUserSignUpRequestBuilder().
		SetEmail(fmt.Sprintf("%v", m["email"])).
		SetPassword(fmt.Sprintf("%v", m["password"])).
		Build()

	signUpInteractor := interactors.NewUserSignUpInteractor(signUpRequest.Email, signUpRequest.Password)
	userSignedUpDto, err1 := signUpInteractor.Execute()
	if err1 != nil {
		return c.JSON(http.StatusOK, models.NewException(err1.Error()))
	}
	return c.JSON(http.StatusOK, userSignedUpDto)
}