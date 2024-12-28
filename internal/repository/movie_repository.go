package repository

import (
	"database/sql"
	"errors"
	"movie-festival/internal/entity"
)

type MovieRepository interface {
	Create(movie *entity.Movie) (int64, error)
	Update(id int, movie *entity.Movie) error
	GetMostViewed() (*entity.Movie, error)
	List(page, pageSize int) ([]*entity.Movie, error)
	Search(query string) ([]*entity.Movie, error)
	TrackViewership(id int) error
}

type movieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) MovieRepository {
	return &movieRepository{db: db}
}

func (repo *movieRepository) Create(movie *entity.Movie) (int64, error) {
	query := `INSERT INTO movies (title, description, duration, artists, genres, watch_url) VALUES (?, ?, ?, ?, ?, ?)`
	result, err := repo.db.Exec(query, movie.Title, movie.Description, movie.Duration, movie.Artists, movie.Genres, movie.WatchURL)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (repo *movieRepository) Update(id int, movie *entity.Movie) error {
	query := `UPDATE movies SET title = ?, description = ?, duration = ?, artists = ?, genres = ?, watch_url = ? WHERE id = ?`
	_, err := repo.db.Exec(query, movie.Title, movie.Description, movie.Duration, movie.Artists, movie.Genres, movie.WatchURL, id)
	return err
}

func (repo *movieRepository) GetMostViewed() (*entity.Movie, error) {
	query := `SELECT id, title, description, duration, artists, genres, watch_url, views FROM movies ORDER BY views DESC LIMIT 1`
	var movie entity.Movie

	err := repo.db.QueryRow(query).Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Duration, &movie.Artists, &movie.Genres, &movie.WatchURL, &movie.Views)

	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (repo *movieRepository) List(page, pageSize int) ([]*entity.Movie, error) {
	offset := (page - 1) * pageSize
	query := `SELECT id, title, description, duration, artists, genres, watch_url, views FROM movies LIMIT ? OFFSET ?`
	rows, err := repo.db.Query(query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*entity.Movie
	for rows.Next() {
		var movie entity.Movie
		err = rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Duration, &movie.Artists, &movie.Genres, &movie.WatchURL, &movie.Views)
		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}
	return movies, nil
}

func (repo *movieRepository) Search(query string) ([]*entity.Movie, error) {
	searchQuery := `SELECT id, title, description, duration, artists, genres, watch_url, views FROM movies WHERE title LIKE ? OR description LIKE ? OR artists LIKE ? OR genres LIKE ?`
	rows, err := repo.db.Query(searchQuery, "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*entity.Movie
	for rows.Next() {
		var movie entity.Movie
		err = rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Duration, &movie.Artists, &movie.Genres, &movie.WatchURL, &movie.Views)
		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}
	return movies, nil
}

func (repo *movieRepository) TrackViewership(id int) error {
	var existingID int
	query := `SELECT id FROM movies WHERE id = ?`
	err := repo.db.QueryRow(query, id).Scan(&existingID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("movie not found")
		}
		return err
	}

	query = `UPDATE movies SET views = views + 1 WHERE id = ?`
	_, err = repo.db.Exec(query, id)
	return err
}
