package rest

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/adapter/models"
	models2 "github.com/DuongVu98/passnet-authentication/src/main/go/domain/config"
	"github.com/DuongVu98/passnet-authentication/src/main/go/usecases/interactors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UserInfo(c echo.Context) (err error) {
	uid := c.Param("uid")

	// Wait for getting beans
	beanConfig := <-beanConfigChannel
	getUserInfoInteractor := interactors.NewGetUserInfoInteractor(uid)
	getUserInfoInteractor.UserRepository = beanConfig.(*models2.BeanConfig).UserRepository

	// execute the interactor
	userDto, getUserError := getUserInfoInteractor.Execute()
	if getUserError != nil {
		return c.JSON(http.StatusOK, models.NewException(getUserError.Error()))
	}
	return c.JSON(http.StatusOK, userDto)
}
