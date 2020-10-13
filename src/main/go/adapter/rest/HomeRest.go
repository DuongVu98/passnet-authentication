package rest

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/adapter/gateway"
	"github.com/DuongVu98/passnet-authentication/src/main/go/adapter/models"
	"github.com/labstack/echo/v4"
	"net/http"
)


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

func TestGrpcMessage(echoContext echo.Context) (err error) {
	message := &models.CreateUserMessage{Uid: "hello uid"}
	messageGateway.SendMessage(message)
	return echoContext.String(http.StatusOK, "sent")
}
