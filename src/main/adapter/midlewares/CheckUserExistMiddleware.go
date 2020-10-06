package midlewares

import (
	"github.com/DuongVu98/passnet-authentication/src/main/config/handles"
	"github.com/DuongVu98/passnet-authentication/src/main/domain/channels"
	models2 "github.com/DuongVu98/passnet-authentication/src/main/domain/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"os"
	"reflect"
)

var beanConfigChannel = channels.GetBeanConfigChannel()

func CheckUserExistMiddleware(key string, c echo.Context) (bool, error) {
	user := c.Get(os.Getenv("AUTH_CONTEXT_KEY")).(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	key = claims["uid"].(string)

	handles.Push(reflect.TypeOf(handles.GetBeanRequest{}))
	beanConfig := <-beanConfigChannel
	userRepository := beanConfig.(*models2.BeanConfig).UserRepository

	_, repoErr := userRepository.FindUserByUid(key)
	if repoErr != nil {
		return false, repoErr
	}

	return true, nil
}
