package service

import (
	"testing"

	"github.com/cvzamannow/service-auth/internal/entity"
	"github.com/cvzamannow/service-auth/internal/repository"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FakeUser struct {
	Id       string `faker:"uuid_digit"`
	Email    string `faker:"email"`
	Password string `faker:"password"`
	Role     uint8  `faker:"oneof: 15, 16"`
}

var mockRepo = &repository.AuthRepositoryMock{Mock: &mock.Mock{}}
var mockService = AuthServiceMock{Repository: mockRepo}

func TestFailSignUp(t *testing.T) {
	// Generate random data
	newUser := FakeUser{}
	errFaker := faker.FakeData(&newUser)

	if errFaker != nil {
		t.Error(errFaker)
	}

	arg := entity.UserEntity{
		Id:       newUser.Id,
		Email:    newUser.Email,
		Password: newUser.Password,
		Role:     int32(newUser.Role),
	}

	// If SaveUser() called, the fake data is should be error and faile to insert record into database
	mockRepo.Mock.On("SaveUser", &arg).Return(nil)
	_, err := mockService.SignUp(arg)

	// Case: There should be any error
	assert.Error(t, err)
}
