package interactors

import (
	"errors"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/config"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/dto"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/repository"
	"github.com/DuongVu98/passnet-authentication/src/main/go/usecases/services"
)

type UserLoginInteractor struct {
	EmailOrDisplayName string
	Password           string
	UserRepository     repository.IUserRepository
}

func NewUserLoginInteractor(emailOrDisplayName string, password string) *UserLoginInteractor {
	beanConfig := config.GetSingletonFactory().Get("beanConfig").(*config.BeanConfig)
	return &UserLoginInteractor{EmailOrDisplayName: emailOrDisplayName, Password: password, UserRepository: beanConfig.UserRepository}
}

func (interactor *UserLoginInteractor) Execute() (*dto.UserTokenDto, error) {
	userEntity, err := interactor.UserRepository.FindUserByEmailOrDisplayName(interactor.EmailOrDisplayName)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}
	isValid := false
	for _, char := range services.Charset {
		passwordToCheck := services.Encrypt(interactor.Password + userEntity.Auth.Salt + string(char))
		if passwordToCheck == userEntity.Auth.Password {
			isValid = true
			break
		}
	}

	if isValid {
		userDto := dto.NewUserDtoBuilder().
			SetUid(userEntity.ID.Hex()).
			SetEmail(userEntity.Email).
			SetDisplayName(userEntity.DisplayName).
			Build()

		// Gen token
		signedToken, jwtErr := services.GenerateJwt(userDto)
		if jwtErr != nil {
			return nil, jwtErr
		}
		userTokenDto := dto.NewUserTokenDtoBuilder().
			SetUserDto(userDto).
			SetToken(signedToken).
			Build()

		return userTokenDto, nil
	} else {
		return nil, errors.New("invalid username or password")
	}
}
