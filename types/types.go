package types

import "github.com/golang-jwt/jwt/v4"

type User struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
}

type Post struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	UserID   string `json:"user_id"`
	PostedAt int64  `json:"posted_at"`
	Liked    bool   `json:"liked"`
}

type Like struct {
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}

type Comment struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	PostID   string `json:"post_id"`
	Text     string `json:"text"`
	PostedAt int64  `json:"posted_at"`
}

type JWTObject struct {
	jwt.RegisteredClaims
	UserID string
}
