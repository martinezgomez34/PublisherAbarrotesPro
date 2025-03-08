package application

import (
	"publisher/src/user/domain"
)

type UserUseCase struct {
	Repo domain.IUserRepository
}

func NewUserUseCase(repo domain.IUserRepository) *UserUseCase {
	return &UserUseCase{Repo: repo}
}

func (uc *UserUseCase) CreateUser(message domain.Message) error {
	return uc.Repo.PublishUser(message)
}