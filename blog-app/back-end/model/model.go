package model

import "time"

// User model
type User struct {
	ID             uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username       string `json:"username" gorm:"uniqueIndex"`
	HashedPassword string `json:"-"` // do not expose in JSON
	SessionToken   string `json:"session_token" gorm:"uniqueIndex"`
	CSRFToken      string `json:"csrf_token"`

	Posts []Post `json:"posts" gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Post model
type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  uint      `json:"author_id"` // Foreign key to User.ID
	Author    User      `json:"author" gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
