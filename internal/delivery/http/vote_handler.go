package http

import (
	"encoding/json"
	"movie-festival/internal/usecase"
	"net/http"
)

type VoteHandler struct {
	usecase usecase.VoteUsecase
}

func NewVoteHandler(voteUsecase usecase.VoteUsecase) *VoteHandler {
	return &VoteHandler{usecase: voteUsecase}
}

func (h *VoteHandler) CreateVote(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	movieIDStr := r.URL.Query().Get("movie_id")

	err := h.usecase.CreateVote(userIDStr, movieIDStr)
	if err != nil {
		http.Error(w, "Failed to create vote", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Vote created successfully"})
}

func (h *VoteHandler) RemoveVote(w http.ResponseWriter, r *http.Request) {
	voteIDStr := r.URL.Query().Get("id")

	err := h.usecase.RemoveVote(voteIDStr)
	if err != nil {
		http.Error(w, "Failed to remove vote", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Vote removed successfully"})
}

func (h *VoteHandler) GetUserVotes(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")

	votes, err := h.usecase.GetUserVotes(userIDStr)
	if err != nil {
		http.Error(w, "Failed to retrieve user votes", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(votes)
}

func (h *VoteHandler) GetMostVotedMovie(w http.ResponseWriter, r *http.Request) {
	movie, err := h.usecase.GetMostVotedMovie()
	if err != nil {
		http.Error(w, "Failed to retrieve most voted movie", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

func (h *VoteHandler) GetMostViewedGenre(w http.ResponseWriter, r *http.Request) {
	genre, err := h.usecase.GetMostViewedGenre()
	if err != nil {
		http.Error(w, "Failed to retrieve most viewed genre", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"genre": genre})
}
