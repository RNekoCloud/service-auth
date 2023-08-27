package service

import (
	"github.com/cvzamannow/service-auth/internal/entity"
	"github.com/cvzamannow/service-auth/internal/repository"
)

type AuthServiceMock struct {
	Repository repository.AuthRepository
}

func (service AuthServiceMock) SignUp(user entity.UserEntity) (string, error) {

	msg, err := service.Repository.SaveUser(&entity.UserEntity{
		Id:       user.Id,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	})

	if err != nil {
		return "", err
	}

	return msg, nil
}
