package googleauth

import (
	"cli-todo/internal/features/user"
)

type UserFinder interface {
	FindUserBySubAndIssuer(sub, issuer string) (*user.Account, error)
}
type UserCreator interface {
	CreateNewAccount(user *user.Account) error
}

type Service struct {
	UserFinder
	UserCreator
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{nil, nil, repo}
}

func (service *Service) FindUserBySubAndIssuer(sub, issuer string) (*user.Account, error) {
	return nil, nil
}
