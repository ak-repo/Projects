package model

import (
	"time"
)

// Post model

type Post struct {
	PostID     string    `json:"postid"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Auther     string    `json:"author"`
	Date       time.Time `json:"createdAt"`
	LastUpdate time.Time `json:"lastupdated"`
}

// // post update model

// type UpdatePost struct {
// 	Title      string    `json:"title"`
// 	Content    string    `json:"content"`
// 	LastUpdate time.Time `json:"lastupdated"`
// }

// user DB
type User struct {
	Username       string
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

type UserRegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
