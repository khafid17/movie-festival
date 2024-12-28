package usecase_test

import (
	"errors"
	"movie-festival/internal/entity"
	"movie-festival/internal/repository"
	"movie-festival/internal/usecase"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserUsecase_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	uc := usecase.NewUserUsecase(mockRepo)

	user := &entity.User{
		Username: "testuser",
		Password: "password",
	}

	mockRepo.EXPECT().Create(user).Return(int64(1), nil).Times(1)

	id, err := uc.Register(user)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)
}

func TestUserUsecase_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	uc := usecase.NewUserUsecase(mockRepo)

	user := &entity.User{
		ID:       1,
		Username: "testuser",
		Password: "password",
	}

	mockRepo.EXPECT().FindByUsername(user.Username).Return(user, nil).Times(1)

	loggedInUser, err := uc.Login(user.Username, user.Password)
	assert.NoError(t, err)
	assert.Equal(t, user, loggedInUser)
}

func TestUserUsecase_Login_InvalidCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	uc := usecase.NewUserUsecase(mockRepo)

	mockRepo.EXPECT().FindByUsername("nonexistent").Return(nil, errors.New("user not found")).Times(1)

	_, err := uc.Login("nonexistent", "wrongpassword")
	assert.Error(t, err)
}

func TestUserUsecase_Logout(t *testing.T) {
	uc := usecase.NewUserUsecase(nil) // Logout tidak membutuhkan repository

	err := uc.Logout()
	assert.NoError(t, err)
}
