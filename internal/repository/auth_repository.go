package repository

import (
	"github.com/cvzamannow/service-auth/internal/entity"
)

type AuthRepository interface {
	SaveUser(user *entity.UserEntity) (string, error)
	CheckUser(email string) (*entity.UserEntity, error)
}
