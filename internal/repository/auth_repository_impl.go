package repository

import (
	"context"

	db "github.com/cvzamannow/service-auth/internal/db/sqlc"
	"github.com/cvzamannow/service-auth/internal/entity"
)

func NewAuthRepository(q *db.Queries) AuthRepository {
	return &authRepositoryImpl{
		Q: q,
	}
}

type authRepositoryImpl struct {
	Q *db.Queries
}

func (repository *authRepositoryImpl) SaveUser(user *entity.UserEntity) (string, error) {
	newUser := db.CreateUserParams{
		ID:       user.Id,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}

	err := repository.Q.CreateUser(context.Background(), newUser)

	if err != nil {
		return "", err
	}

	return "Success", nil
}

func (repository *authRepositoryImpl) CheckUser(email string) (*entity.UserEntity, error) {
	result, err := repository.Q.FindUser(context.Background(), email)

	if err != nil {
		return nil, err
	}

	return &entity.UserEntity{
		Id:       result.ID,
		Email:    result.Email,
		Password: result.Password,
		Role:     result.Role,
	}, nil
}
