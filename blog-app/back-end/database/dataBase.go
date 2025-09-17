package database

import (
	"gin-blog-app/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	var err error

	cst := "user=ak dbname=blog password=4455@mint sslmode=disable"

	DB, err = gorm.Open(postgres.Open(cst), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	if err := DB.AutoMigrate(&model.User{}, &model.Post{}); err != nil {
		log.Fatal("Database migration are failed")
	}
	log.Println("Connected into DB ")

}

