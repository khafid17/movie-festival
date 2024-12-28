package usecase

import (
	"movie-festival/internal/entity"
	"movie-festival/internal/repository"
)

type VoteUsecase interface {
	CreateVote(userID, movieID string) error
	RemoveVote(voteID string) error
	GetUserVotes(userID string) ([]entity.Vote, error)
	GetMostVotedMovie() (*entity.Movie, error)
	GetMostViewedGenre() (string, error)
}

type VoteUsecaseImpl struct {
	repo repository.VoteRepository
}

func NewVoteUsecase(repo repository.VoteRepository) *VoteUsecaseImpl {
	return &VoteUsecaseImpl{repo: repo}
}

func (uc *VoteUsecaseImpl) CreateVote(userID, movieID string) error {
	return uc.repo.Create(userID, movieID)
}

func (uc *VoteUsecaseImpl) RemoveVote(voteID string) error {
	return uc.repo.Remove(voteID)
}

func (uc *VoteUsecaseImpl) GetUserVotes(userID string) ([]entity.Vote, error) {
	return uc.repo.GetUserVotes(userID)
}

func (uc *VoteUsecaseImpl) GetMostVotedMovie() (*entity.Movie, error) {
	return uc.repo.GetMostVotedMovie()
}

func (uc *VoteUsecaseImpl) GetMostViewedGenre() (string, error) {
	return uc.repo.GetMostViewedGenre()
}
