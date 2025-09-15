package database

import (
	"database/sql"
	"gin-blog-app/model"
	"log"

	_ "github.com/lib/pq"
)

// user db

//map[string]model.User{}

var UsersDB = []model.User{}

// post db
var PostList = []model.Post{}

var DB *sql.DB

// init db

func InitDB() {

	var err error

	connStr := "user=ak dbname=blogs password=4455@mint sslmode=disable"

	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("DB connection failed: ,", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("DB ping failed: ", err)
	}
	log.Println("✅ Connected to Postgres")

}

// _, err := database.DB.Exec(
// 	"INSERT INTO users (username, hashed_password) VALUES ($1, $2)",
// 	user.Username, user.HashedPassword,
// )
// if err != nil {
// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "DB insert failed"})
// 	return
// }
