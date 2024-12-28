package entity

type Vote struct {
	ID        int    `json:"id"`
	UserID    string `json:"user_id"`
	MovieID   string `json:"movie_id"`
	Voted     bool   `json:"voted"`
	CreatedAt string `json:"created_at"`
}
