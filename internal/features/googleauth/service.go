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

type AuthService struct {
	UserFinder
	UserCreator
	repo Repository
}

func NewService(repo Repository) *AuthService {
	return &AuthService{nil, nil, repo}
}

func (service *AuthService) CreateNewSessionToken(userInfo OauthUserInfo) (*user.Account, error) {
	user, err := service.UserFinder.FindUserBySubAndIssuer(userInfo.Sub, string(userInfo.Issuer))

	return nil, nil
}
