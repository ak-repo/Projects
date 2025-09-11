package main

import (
	"gin-blog-app/controler"
	"gin-blog-app/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	//multiplexer
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}), gin.Recovery(), middleware.LoggerMiddleware())

	//public routes
	router.GET("/", controler.HomeHandler)
	router.GET("/posts", controler.PostPageHandler)
	router.POST("/register", controler.RegistrationHandler)
	router.POST("/login", controler.LoginHandler)
	router.POST("/logout", controler.LogoutHandler)

	// private routes

	private := router.Group("/user")
	private.Use(middleware.RequireAuth())
	{

		// post creation, updation, delete routes
		private.POST("/create", controler.CreatePostHandler)
		private.PATCH("/update/:postid", middleware.RequireOwner(), controler.UpdatePostHandler)
		private.DELETE("/delete/:postid", middleware.RequireOwner(), controler.DeletePostHandler)

	}

	//server
	router.Run()

}
