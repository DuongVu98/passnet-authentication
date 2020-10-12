package rest

import (
	"context"
	"github.com/DuongVu98/passnet-authentication/src/main/go/adapter/gateway"
	"github.com/DuongVu98/passnet-authentication/src/main/go/adapter/models"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/channels"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/config"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/dto"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

//var testConfig = config.GetAppConfig()
//var firebaseApp = testConfig.FirebaseApp

var appConfigChannel = channels.GetAppConfigChannel()
var messageGateway = gateway.NewMessageGateway()

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func HomePage(c echo.Context) (err error) {
	return c.String(http.StatusOK, "This is home page\n")
}

func JsonResponseSample(c echo.Context) (err error) {
	var message = "Response from server"
	return c.JSON(http.StatusOK, message)
}

func UserRetrieve(echoContext echo.Context) (err error) {
	appConfig := <-appConfigChannel
	firebaseApp := appConfig.(config.AppConfig).FirebaseApp

	client, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	u, err := client.GetUser(context.Background(), "PdwwmEelwoXV97msTviL2rP0LQz2")
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", "PdwwmEelwoXV97msTviL2rP0LQz2", err)
	}
	log.Println("Successfully fetched user data")
	userDto := dto.NewUserDtoBuilder().
		SetUid(u.UID).
		SetDisplayName(u.DisplayName).
		SetEmail(u.Email).
		Build()

	return echoContext.JSON(http.StatusOK, userDto)
}

func TestGrpcMessage(echoContext echo.Context) (err error) {
	message := &models.CreateUserMessage{Uid: "hello uid"}
	messageGateway.SendMessage(message)
	return echoContext.String(http.StatusOK, "sent")
}
