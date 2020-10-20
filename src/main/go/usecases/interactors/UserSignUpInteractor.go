package interactors

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/config"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/dto"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/entity"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/repository"
	"github.com/DuongVu98/passnet-authentication/src/main/go/usecases/services"
	"log"
)

type UserSignUpInteractor struct {
	Email          string
	Password       string
	UserRepository repository.IUserRepository
}

func NewUserSignUpInteractor(email string, password string) *UserSignUpInteractor {
	beanConfig := config.GetSingletonFactory().Get("beanConfig").(*config.BeanConfig)
	return &UserSignUpInteractor{Email: email, Password: password, UserRepository: beanConfig.UserRepository}
}

func (interactor *UserSignUpInteractor) Execute() (*dto.UserDto, error) {

	newSalt := services.RandomSalt()
	newPepper := services.RandomCharacter()
	passWordToSave := services.Encrypt(interactor.Password + newSalt + newPepper)
	newUserAuth := entity.NewUserAuthentication(passWordToSave, newSalt)

	newUser := entity.NewUserEntityBuilder().
		SetEmail(interactor.Email).
		SetAuthentication(newUserAuth).
		Build()

	log.Printf("salt: %s, pepper: %s, encrypt: %s", newSalt, newPepper, passWordToSave)

	userSignedUpEntity, err := interactor.UserRepository.InsertUser(newUser)
	if err != nil {
		return nil, err
	}

	log.Printf("user after insert: %v", userSignedUpEntity)
	userSignedUpDto := dto.NewUserDtoBuilder().
		SetEmail(userSignedUpEntity.Email).
		SetDisplayName(userSignedUpEntity.DisplayName).
		SetUid(userSignedUpEntity.ID.Hex()).
		Build()
	return userSignedUpDto, nil
}
