package services

import (
	"github.com/DuongVu98/passnet-authentication/src/main/domain/dto"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateJwt(userDto *dto.UserDto) (string, error) {
	// Gen token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = userDto.Uid
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	// Generate encoded token and send it as response.
	signedToken, jwtErr := token.SignedString([]byte("secret"))
	if jwtErr != nil {
		return "", jwtErr
	}
	return signedToken, nil
}