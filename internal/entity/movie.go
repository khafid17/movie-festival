package entity

type Movie struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Artists     string `json:"artists"`
	Genres      string `json:"genres"`
	WatchURL    string `json:"watch_url"`
	Views       int    `json:"views"`
}