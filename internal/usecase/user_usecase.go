package usecase

import (
	"movie-festival/internal/entity"
	"movie-festival/internal/repository"
)

type UserUsecase interface {
	Register(user *entity.User) (int64, error)
	Login(username, password string) (*entity.User, error)
	Logout() error
}

type UserUsecaseImpl struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecaseImpl {
	return &UserUsecaseImpl{repo: repo}
}

func (uc *UserUsecaseImpl) Register(user *entity.User) (int64, error) {
	return uc.repo.Create(user)
}

func (uc *UserUsecaseImpl) Login(username, password string) (*entity.User, error) {
	return uc.repo.FindByUsername(username)
}

func (uc *UserUsecaseImpl) Logout() error {
	return nil
}
