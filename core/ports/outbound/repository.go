package outbound

import (
	"github.com/roku-on-it/golang-search/core/domain"
)

type UserRepository interface {
	FindUserByID(id string) (*domain.User, error)
	FindUserByUsername(u string) (*domain.User, error)
	CreateUser(input domain.CreateUserInput) (*domain.User, error)
}
