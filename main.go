package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"movie-festival/internal/config"
	httpDelivery "movie-festival/internal/delivery/http"
	"movie-festival/internal/repository"
	"movie-festival/internal/usecase"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database", err)
	}
	defer db.Close()

	movieRepo := repository.NewMovieRepository(db)
	movieUsecase := usecase.NewMovieUsecase(movieRepo)
	movieHandler := httpDelivery.NewMovieHandler(movieUsecase)

	voteRepo := repository.NewVoteRepository(db)
    voteUsecase := usecase.NewVoteUsecase(voteRepo)
    voteHandler := httpDelivery.NewVoteHandler(voteUsecase)

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)	
	userHandler := httpDelivery.NewUserHandler(userUsecase)

	
	r := mux.NewRouter()
	
	// movies 
	r.HandleFunc("/movies", movieHandler.CreateMovie).Methods("POST")
	r.HandleFunc("/movies", movieHandler.ListMovies).Methods("GET")
	r.HandleFunc("/movies/search", movieHandler.SearchMovies).Methods("GET")
	r.HandleFunc("/movies/{id}/view", movieHandler.TrackMovieViewership).Methods("POST")
	r.HandleFunc("/movies/{id}", movieHandler.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/most-viewed", movieHandler.GetMostViewed).Methods("GET")
	
	// votes 
	r.HandleFunc("/votes", voteHandler.CreateVote).Methods("POST")
	r.HandleFunc("/unvotes", voteHandler.RemoveVote).Methods("DELETE")
	r.HandleFunc("/votes/list", voteHandler.GetUserVotes).Methods("GET")
	r.HandleFunc("/votes/most-voted", voteHandler.GetMostVotedMovie).Methods("GET")
	r.HandleFunc("/votes/most-viewed-genre", voteHandler.GetMostViewedGenre).Methods("GET")

	// users 
	r.HandleFunc("/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/logout", userHandler.Logout).Methods("POST")

	log.Fatal(http.ListenAndServe(":8181", r))
}
