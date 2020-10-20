package interactors

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/config"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/dto"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/repository"
)

type GetUserInfoInteractor struct {
	RequestUid     string
	UserRepository repository.IUserRepository
}

func NewGetUserInfoInteractor(requestUid string) *GetUserInfoInteractor {
	beanConfig := config.GetSingletonFactory().Get("beanConfig").(*config.BeanConfig)
	return &GetUserInfoInteractor{RequestUid: requestUid, UserRepository: beanConfig.UserRepository}
}

func (interactor *GetUserInfoInteractor) Execute() (*dto.UserDto, error) {
	userEntity, err := interactor.UserRepository.FindUserByUid(interactor.RequestUid)
	if err != nil {
		return nil, err
	}
	userDto := dto.NewUserDtoBuilder().
		SetUid(userEntity.ID.Hex()).
		SetDisplayName(userEntity.DisplayName).
		SetEmail(userEntity.Email).
		Build()
	return userDto, nil
}
