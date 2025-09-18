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

	// check the username already exist in DB
	var existingUser model.User
	if err := database.DB.Where("username = ?", newUser.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username alreay exists"})
		return

	}

	//password hashing
	hashedPassword, err := utils.GenerateHashedPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "password hashing failed"})
		return
	}

	// add user into dp
	user := model.User{
		Username:       newUser.Username,
		HashedPassword: hashedPassword,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true})

}

// Login
func LoginHandler(c *gin.Context) {

	var loginUser model.UserLoginInput
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	// find user by username
	if err := database.DB.Where("username=?", loginUser.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
		return
	}

	// Verify password
	if !utils.CompareHashAndPassword(loginUser.Password, user.HashedPassword) {
		c.JSON(http.StatusConflict, gin.H{"error": "Invalid password"})
		return
	}

	//token generation
	sessionToken := utils.GenerateToken(32)
	csrfToken := utils.GenerateToken(32)

	// Update tokens on DB
	if err := database.DB.Model(&user).Updates(map[string]interface{}{
		"session_token": sessionToken,
		"csrf_token":    csrfToken,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tokens"})
		return

	}

	//cookie
	c.SetCookie("session_token", sessionToken, 3600*24, "/", "", false, true)
	c.SetCookie("csrf_token", csrfToken, 3600*24, "/", "", false, false)
	c.JSON(http.StatusOK, gin.H{"ID": user.ID, "username": user.Username})

}

// Logout
func LogoutHandler(c *gin.Context) {

	sToken, err := c.Cookie("session_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no session found"})
		return
	}

	// update in db
	if err := database.DB.Model(&model.User{}).Where("session_token = ?", sToken).Updates(map[string]interface{}{
		"session_token": "",
		"csrf_token":    "",
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to logout"})
		return
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
		c.JSON(http.StatusConflict, gin.H{"error": "Invalid formate"})
		return
	}

	autherID := c.MustGet("userID").(uint)

	newPost.CreatedAt = time.Now()
	newPost.AuthorID = autherID

	if err := database.DB.Create(&newPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cound upload right now"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": newPost})

}

// delete
func DeletePostHandler(c *gin.Context) {
	id := c.Param("postid")

	//delete from db
	if err := database.DB.Delete(&model.Post{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could't delete from server"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})

}

// update post
func UpdatePostHandler(c *gin.Context) {
	id := c.Param("postid")

	// check the post in DB
	var existingPost model.Post
	if err := database.DB.First(&existingPost, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return

	}

	var updateData struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Invalide request body"})
		return
	}

	//update old post with new
	if err := database.DB.Model(&existingPost).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"updated": false})

	}

	c.JSON(http.StatusOK, gin.H{"updated": true, "post": existingPost})

}

// post page handler
func PostPageHandler(c *gin.Context) {
	var posts []model.Post
	if err := database.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found anything"})
		return
	}
	// var response

	c.JSON(http.StatusOK, gin.H{"posts": posts})

}

func GetPostByID(c *gin.Context) {

	id := c.Param("id")

	var post model.Post

	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	//find the auther
	var auther model.User
	if err := database.DB.First(&auther, post.AuthorID).Error; err != nil {
		log.Println(post.AuthorID)
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return

	}

	//response
	response := model.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		Author:    auther.Username,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
	c.JSON(http.StatusOK, gin.H{"post": response})

}
