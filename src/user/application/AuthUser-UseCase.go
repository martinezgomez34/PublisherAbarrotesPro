package application

import (
	"publisher/src/user/domain"
)

type AuthUseCase struct {
	Repo domain.IUserRepository
}

func NewAuthUseCase(repo domain.IUserRepository) *AuthUseCase {
	return &AuthUseCase{Repo: repo}
}

func (uc *AuthUseCase) Login(message domain.Message) error {
	return uc.Repo.PublishLogin(message)
}