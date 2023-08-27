package repository

import (
	"errors"

	"github.com/cvzamannow/service-auth/internal/entity"
	"github.com/stretchr/testify/mock"
)

type AuthRepositoryMock struct {
	Mock *mock.Mock
}

func (repository *AuthRepositoryMock) SaveUser(user *entity.UserEntity) (string, error) {
	args := repository.Mock.Called(user)
	if args.Get(0) == nil {
		return "", errors.New("Failed to insert database")
	}

	return args.String(0), nil

}

func (repository *AuthRepositoryMock) CheckUser(email string) (*entity.UserEntity, error) {
	args := repository.Mock.Called(email)

	if args.Get(0) == nil {
		return nil, errors.New("no argument was found in service")
	} else {
		user := args.Get(0).(entity.UserEntity)
		return &user, nil
	}
}
