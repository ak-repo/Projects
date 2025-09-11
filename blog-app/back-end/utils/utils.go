package utils

import (
	"crypto/rand"
	"encoding/base64"
	"gin-blog-app/database"
	"gin-blog-app/model"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// id generator function
func GenerateID() string {
	return uuid.New().String()
}

// password hasing algorithum
func GenerateHashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// compare hashed password
func CompareHashAndPassword(password, hased string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hased))
	return err == nil
}

// token generation
func GenerateToken(length int) string {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Token generate Failed: %v", err)
	}
	return base64.URLEncoding.EncodeToString(bytes)

}

//----------------------------------USER-----------------------------------------

// db checker is in DB
func UserFinder(username string) (model.User, bool) {

	var user model.User

	for _, u := range database.UsersDB {
		if u.Username == username {
			return u, true
		}
	}

	return user, false
}

// update user info
func UserUpdater(user model.User) bool {
	for i, u := range database.UsersDB {
		if u.Username == user.Username {
			database.UsersDB[i] = user
			return true
		}
	}

	return false

}

//------------------------------------------POST---DB----utils--------------------------

