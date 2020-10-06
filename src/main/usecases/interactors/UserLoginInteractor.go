package interactors

import (
	"errors"
	"github.com/DuongVu98/passnet-authentication/src/main/domain/dto"
	"github.com/DuongVu98/passnet-authentication/src/main/domain/repository"
	"github.com/DuongVu98/passnet-authentication/src/main/usecases/services"
)

type UserLoginInteractor struct {
	EmailOrDisplayName string
	Password           string
	UserRepository     repository.IUserRepository
}

func NewUserLoginInteractor(emailOrDisplayName string, password string) *UserLoginInteractor {
	return &UserLoginInteractor{EmailOrDisplayName: emailOrDisplayName, Password: password}
}

func (interactor *UserLoginInteractor) Execute() (*dto.UserDto, error) {
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
		return userDto, nil
	} else {
		return nil, errors.New("invalid username or password")
	}
}
