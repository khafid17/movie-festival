package usecase

import (
	"movie-festival/internal/entity"
	"movie-festival/internal/repository"
)

type MovieUsecase interface {
	CreateMovie(movie *entity.Movie) (int64, error)
	UpdateMovie(id int, movie *entity.Movie) error
	GetMostViewed() (*entity.Movie, error)
	ListMovies(page, pageSize int) ([]*entity.Movie, error)
	SearchMovies(query string) ([]*entity.Movie, error)
	TrackMovieViewership(id int) error
}

type movieUsecase struct {
	repo repository.MovieRepository
}

func NewMovieUsecase(repo repository.MovieRepository) MovieUsecase {
	return &movieUsecase{repo: repo}
}

func (uc *movieUsecase) CreateMovie(movie *entity.Movie) (int64, error) {
	return uc.repo.Create(movie)
}

func (uc *movieUsecase) UpdateMovie(id int, movie *entity.Movie) error {
	return uc.repo.Update(id, movie)
}

func (uc *movieUsecase) GetMostViewed() (*entity.Movie, error) {
	return uc.repo.GetMostViewed()
}

func (uc *movieUsecase) ListMovies(page, pageSize int) ([]*entity.Movie, error) {
	return uc.repo.List(page, pageSize)
}

func (uc *movieUsecase) SearchMovies(query string) ([]*entity.Movie, error) {
	return uc.repo.Search(query)
}

func (uc *movieUsecase) TrackMovieViewership(id int) error {
	return uc.repo.TrackViewership(id)
}
