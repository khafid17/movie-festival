package usecase_test

import (
	"movie-festival/internal/entity"
	"movie-festival/internal/repository"
	"movie-festival/internal/usecase"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockMovieRepository(ctrl)
	movieUsecase := usecase.NewMovieUsecase(mockRepo)

	movie := &entity.Movie{
		Title:       "Test Movie",
		Description: "Description",
		Duration:    120,
		Artists:     "Artist 1, Artist 2",
		Genres:      "Action, Drama",
		WatchURL:    "http://example.com",
	}

	mockRepo.EXPECT().Create(movie).Return(int64(1), nil)

	id, err := movieUsecase.CreateMovie(movie)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)
}

func TestGetMostViewed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockMovieRepository(ctrl)
	movieUsecase := usecase.NewMovieUsecase(mockRepo)

	expectedMovie := &entity.Movie{
		ID:          1,
		Title:       "Most Viewed Movie",
		Description: "Popular movie",
		Duration:    120,
		Artists:     "Actor 1, Actor 2",
		Genres:      "Action",
		WatchURL:    "http://example.com",
		Views:       1000,
	}

	mockRepo.EXPECT().GetMostViewed().Return(expectedMovie, nil)

	movie, err := movieUsecase.GetMostViewed()
	assert.NoError(t, err)
	assert.Equal(t, expectedMovie, movie)
}

func TestListMovies(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockMovieRepository(ctrl)
	movieUsecase := usecase.NewMovieUsecase(mockRepo)

	expectedMovies := []*entity.Movie{
		{ID: 1, Title: "Movie 1"},
		{ID: 2, Title: "Movie 2"},
	}

	mockRepo.EXPECT().List(1, 10).Return(expectedMovies, nil)

	movies, err := movieUsecase.ListMovies(1, 10)
	assert.NoError(t, err)
	assert.Equal(t, expectedMovies, movies)
}

func TestSearchMovies(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockMovieRepository(ctrl)
	movieUsecase := usecase.NewMovieUsecase(mockRepo)

	query := "Test"
	expectedMovies := []*entity.Movie{
		{ID: 1, Title: "Test Movie 1"},
		{ID: 2, Title: "Test Movie 2"},
	}

	mockRepo.EXPECT().Search(query).Return(expectedMovies, nil)

	movies, err := movieUsecase.SearchMovies(query)
	assert.NoError(t, err)
	assert.Equal(t, expectedMovies, movies)
}

func TestTrackMovieViewership(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockMovieRepository(ctrl)
	movieUsecase := usecase.NewMovieUsecase(mockRepo)

	mockRepo.EXPECT().TrackViewership(1).Return(nil)

	err := movieUsecase.TrackMovieViewership(1)
	assert.NoError(t, err)
}

func TestUpdateMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockMovieRepository(ctrl)
	movieUsecase := usecase.NewMovieUsecase(mockRepo)

	movie := &entity.Movie{
		Title:       "Updated Movie",
		Description: "Updated Description",
		Duration:    130,
		Artists:     "Updated Artist",
		Genres:      "Updated Genre",
		WatchURL:    "http://updated.com",
	}

	mockRepo.EXPECT().Update(1, movie).Return(nil)

	err := movieUsecase.UpdateMovie(1, movie)
	assert.NoError(t, err)
}
