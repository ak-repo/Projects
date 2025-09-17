package middleware

import (
	"gin-blog-app/database"
	"gin-blog-app/model"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Authentication middleware
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sToken, err := c.Cookie("session_token")
		if err != nil || sToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "please login first"})
			c.Abort()
			return
		}

		// check if user exists with this session
		var currentUser model.User
		if err := database.DB.Where("session_token = ?", sToken).First(&currentUser).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
			c.Abort()
			return
		}

		// store userID in context
		c.Set("userID", currentUser.ID)
		c.Next()
	}
}

// Authorization middleware (for delete/update posts)
func RequireOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get logged-in user ID from context
		userID := c.MustGet("userID").(uint)

		// get postID from URL and convert to uint
		postIDParam := c.Param("postid")
		postID, err := strconv.ParseUint(postIDParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
			c.Abort()
			return
		}

		// fetch post
		var post model.Post
		if err := database.DB.First(&post, postID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
			c.Abort()
			return
		}

		// check ownership
		if post.AuthorID != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": "not allowed to modify this post"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Logger middleware
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		log.Printf("[LOGGER] %s %s | %d | %v",
			c.Request.Method, c.Request.URL.Path, status, latency)
	}
}
