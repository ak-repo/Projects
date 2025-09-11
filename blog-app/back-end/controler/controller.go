package controler

import (
	"gin-blog-app/database"
	"gin-blog-app/model"
	"gin-blog-app/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// home page
func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Home"})

}

// post page handler
func PostPageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, database.PostList)

}

//---------------------------AUth----------------------------------------------

// registration
func RegistrationHandler(c *gin.Context) {

	var newUser model.UserRegisterInput

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validation
	if len(newUser.Username) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username must be at least 6 characters"})
		return
	}

	if len(newUser.Password) < 9 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 9 characters"})
		return
	}

	//check already exist
	if _, ok := utils.UserFinder(newUser.Username); ok {
		c.JSON(http.StatusConflict, gin.H{"error": ok})
		return

	}

	//password hashing
	hashedPassword, _ := utils.GenerateHashedPassword(newUser.Password)

	// add user into dp
	user := model.User{
		Username:       newUser.Username,
		HashedPassword: hashedPassword,
	}
	database.UsersDB = append(database.UsersDB, user)

	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})

}

// Login
func LoginHandler(c *gin.Context) {

	var loginUser model.UserLoginInput
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, ok := utils.UserFinder(loginUser.Username)

	//verification:= password, username
	if !ok || utils.CompareHashAndPassword(loginUser.Password, user.HashedPassword) {
		c.JSON(http.StatusConflict, gin.H{"error": "Invalid password or username"})
		return
	}

	//token generation
	sessionToken := utils.GenerateToken(32)
	csrfToken := utils.GenerateToken(32)

	//cookie

	c.SetCookie("session_token", sessionToken, 3600*24, "/", "", false, true)
	c.SetCookie("csrf_token", csrfToken, 3600*24, "/", "", false, false)

	// update on db
	user.CSRFToken = csrfToken
	user.SessionToken = sessionToken

	utils.UserUpdater(user)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})

}

// Logout
func LogoutHandler(c *gin.Context) {

	sToken, err := c.Cookie("session_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no session found"})
		return
	}

	// check in DB
	for _, user := range database.UsersDB {

		if user.SessionToken == sToken {
			user.CSRFToken = ""
			user.SessionToken = ""
			utils.UserUpdater(user)
			break
		}
	}

	c.SetCookie("session_token", "", -1, "/", "", true, true)
	c.SetCookie("csrf_token", "", -1, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successfull"})

}

//----------------------------------POST------------------------------------------

// creation
func CreatePostHandler(c *gin.Context) {
	var newPost model.Post

	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "not in proper form"})
		return
	}

	auther := c.MustGet("user").(string)
	log.Println(auther)

	//generate id
	id := utils.GenerateID()

	newPost.PostID = id
	newPost.Date = time.Now()
	newPost.Auther = auther
	database.PostList = append(database.PostList, newPost)
	c.JSON(http.StatusOK, gin.H{"success": newPost})

}

// delete
func DeletePostHandler(c *gin.Context) {
	id := c.Param("postid")

	//delete from db

	for index, post := range database.PostList {
		if post.PostID == id {
			database.PostList = append(database.PostList[:index], database.PostList[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"success": true})
			return

		}
	}

	//if not deleted
	c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})

}

// update post
func UpdatePostHandler(c *gin.Context) {
	var updatedPost model.Post
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "post formate error"})
		return
	}

	//update old post with new
	id := c.Param("postid")
	// auther := c.MustGet("user").(string)
	updatedPost.LastUpdate = time.Now()
	updatedPost.PostID = id
	// updatedPost.Auther = auther

	for index, post := range database.PostList {
		if post.PostID == updatedPost.PostID {
			// updatedPost.Date = post.Date
			database.PostList[index] = updatedPost
			c.JSON(http.StatusOK, gin.H{"updated": true})
			return

		}
	}
	c.JSON(http.StatusNotFound, gin.H{"updated": false})

}
