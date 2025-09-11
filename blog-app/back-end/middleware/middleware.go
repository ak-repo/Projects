package middleware

import (
	"gin-blog-app/database"
	"gin-blog-app/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Authentication middleware
func RequireAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		sToken, err := c.Cookie("session_token")
		if err != nil || sToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": " please login first"})
			c.Abort()
			return
		}
		// check the user id
		var currentUser *model.User
		for _, u := range database.UsersDB {
			if u.SessionToken == sToken {
				currentUser = &u
				break
			}
		}

		if currentUser == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
			c.Abort()
			return
		}

		c.Set("user", currentUser.Username)
		c.Next()
	}

}

// Authorization for delete , update

func RequireOwner() gin.HandlerFunc {

	return func(c *gin.Context) {

		user := c.MustGet("user").(string)
		// postID := c.Param("postid")

		for _, post := range database.PostList {
			if post.Auther == user {
				c.Next()
				return

			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "not allowed to modify this post"})
		c.Abort()
	}
}

// logger middleware

func LoggerMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		log.Printf("[LOGGER] %s %s  | %d   |  %v   ",
			c.Request.Method, c.Request.URL.Path, status, latency)

	}
}
