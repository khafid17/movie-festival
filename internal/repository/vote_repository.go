package repository

import (
	"database/sql"
	"movie-festival/internal/entity"
)

type VoteRepository interface {
	Create(userID, movieID string) error
	Remove(voteID string) error
	GetUserVotes(userID string) ([]entity.Vote, error)
	GetMostVotedMovie() (*entity.Movie, error)
	GetMostViewedGenre() (string, error)
}

type VoteRepositoryImpl struct {
	db *sql.DB
}

func NewVoteRepository(db *sql.DB) *VoteRepositoryImpl {
	return &VoteRepositoryImpl{db: db}
}

func (repo *VoteRepositoryImpl) Create(userID, movieID string) error {
	query := `INSERT INTO votes (user_id, movie_id) VALUES (?, ?)`
	_, err := repo.db.Exec(query, userID, movieID)
	return err
}

func (repo *VoteRepositoryImpl) Remove(voteID string) error {
	query := `DELETE FROM votes WHERE id = ?`
	_, err := repo.db.Exec(query, voteID)
	return err
}

func (repo *VoteRepositoryImpl) GetUserVotes(userID string) ([]entity.Vote, error) {
	query := `SELECT id, user_id, movie_id, created_at FROM votes WHERE user_id = ?`
	rows, err := repo.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var votes []entity.Vote
	for rows.Next() {
		vote := entity.Vote{}
		err := rows.Scan(&vote.ID, &vote.UserID, &vote.MovieID, &vote.CreatedAt)
		if err != nil {
			return nil, err
		}
		votes = append(votes, vote)
	}
	return votes, nil
}

func (repo *VoteRepositoryImpl) GetMostVotedMovie() (*entity.Movie, error) {
	query := `SELECT movies.id, movies.title, movies.description, movies.duration, movies.artists, movies.genres, movies.watch_url, COUNT(votes.movie_id) AS views FROM votes JOIN movies ON votes.movie_id = movies.id GROUP BY movies.id ORDER BY views DESC LIMIT 1`
	row := repo.db.QueryRow(query)
	movie := &entity.Movie{}
	err := row.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Duration, &movie.Artists, &movie.Genres, &movie.WatchURL, &movie.Views)

	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (repo *VoteRepositoryImpl) GetMostViewedGenre() (string, error) {
	query := `SELECT genres FROM movies GROUP BY genres ORDER BY COUNT(*) DESC LIMIT 1`
	row := repo.db.QueryRow(query)

	var genres string
	err := row.Scan(&genres)
	if err != nil {
		return "", err
	}
	return genres, nil
}
