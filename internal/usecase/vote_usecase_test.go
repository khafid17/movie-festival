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

func TestVoteUsecase_CreateVote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockVoteRepository(ctrl)
	uc := usecase.NewVoteUsecase(mockRepo)

	userID := "17"
	movieID := "7"

	mockRepo.EXPECT().Create(userID, movieID).Return(nil).Times(1)

	err := uc.CreateVote(userID, movieID)
	assert.NoError(t, err)
}

func TestVoteUsecase_RemoveVote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockVoteRepository(ctrl)
	uc := usecase.NewVoteUsecase(mockRepo)

	voteID := "5"

	mockRepo.EXPECT().Remove(voteID).Return(nil).Times(1)

	err := uc.RemoveVote(voteID)
	assert.NoError(t, err)
}

func TestVoteUsecase_GetUserVotes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockVoteRepository(ctrl)
	uc := usecase.NewVoteUsecase(mockRepo)

	userID := "17"
	expectedVotes := []entity.Vote{
		{ID: 1, UserID: userID, MovieID: "movie123", CreatedAt: "2024-12-29"},
		{ID: 2, UserID: userID, MovieID: "movie456", CreatedAt: "2024-12-30"},
	}

	mockRepo.EXPECT().GetUserVotes(userID).Return(expectedVotes, nil).Times(1)

	votes, err := uc.GetUserVotes(userID)
	assert.NoError(t, err)
	assert.Equal(t, expectedVotes, votes)
}

func TestVoteUsecase_GetUserVotes_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockVoteRepository(ctrl)
	uc := usecase.NewVoteUsecase(mockRepo)

	userID := "17"

	mockRepo.EXPECT().GetUserVotes(userID).Return(nil, errors.New("database error")).Times(1)

	votes, err := uc.GetUserVotes(userID)
	assert.Error(t, err)
	assert.Nil(t, votes)
}

func TestVoteUsecase_GetMostVotedMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockVoteRepository(ctrl)
	uc := usecase.NewVoteUsecase(mockRepo)

	expectedMovie := &entity.Movie{
		ID:          17,
		Title:       "Best Movie",
		Description: "A great movie",
		Duration:    120,
		Artists:     "John Doe, Jane Doe",
		Genres:      "Drama",
		WatchURL:    "http://example.com/movie123",
		Views:       100,
	}

	mockRepo.EXPECT().GetMostVotedMovie().Return(expectedMovie, nil).Times(1)

	movie, err := uc.GetMostVotedMovie()
	assert.NoError(t, err)
	assert.Equal(t, expectedMovie, movie)
}

func TestVoteUsecase_GetMostViewedGenre(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockVoteRepository(ctrl)
	uc := usecase.NewVoteUsecase(mockRepo)

	expectedGenre := "Action"

	mockRepo.EXPECT().GetMostViewedGenre().Return(expectedGenre, nil).Times(1)

	genre, err := uc.GetMostViewedGenre()
	assert.NoError(t, err)
	assert.Equal(t, expectedGenre, genre)
}
