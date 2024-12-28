package http

import (
	"encoding/json"
	"log"
	"movie-festival/internal/entity"
	"movie-festival/internal/usecase"
	"net/http"
	"strconv"
)

type MovieHandler struct {
	Usecase usecase.MovieUsecase
}

func NewMovieHandler(uc usecase.MovieUsecase) *MovieHandler {
	return &MovieHandler{Usecase: uc}
}

func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie entity.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		log.Println("Invalid request payload:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	id, err := h.Usecase.CreateMovie(&movie)
	if err != nil {
		log.Println("Failed to create movie:", err)
		http.Error(w, "Failed to create movie", http.StatusInternalServerError)
		return
	}
	movie.ID = int(id)
	json.NewEncoder(w).Encode(movie)
}

func (h *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	var movie entity.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.Usecase.UpdateMovie(id, &movie); err != nil {
		http.Error(w, "Failed to update movie", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Movie updated successfully"})
}

func (h *MovieHandler) GetMostViewed(w http.ResponseWriter, r *http.Request) {
	movie, err := h.Usecase.GetMostViewed()
	if err != nil {
		http.Error(w, "Failed to retrieve most viewed movie", http.StatusInternalServerError)
		return
	}
    
	json.NewEncoder(w).Encode(movie)
}

func (h *MovieHandler) ListMovies(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	movies, err := h.Usecase.ListMovies(page, pageSize)
	if err != nil {
		http.Error(w, "Failed to list movies", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(movies)
}

func (h *MovieHandler) SearchMovies(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	movies, err := h.Usecase.SearchMovies(query)
	if err != nil {
		http.Error(w, "Failed to search movies", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(movies)
}

func (h *MovieHandler) TrackMovieViewership(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "Invalid movie ID", http.StatusBadRequest)
        return
    }

    id, err := strconv.Atoi(idStr)
    if err != nil {
        log.Printf("Error converting ID: %v", err)
        http.Error(w, "Invalid movie ID", http.StatusBadRequest)
        return
    }

    if err := h.Usecase.TrackMovieViewership(id); err != nil {
        http.Error(w, "Failed to track viewership", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Viewership tracked successfully"})
}
