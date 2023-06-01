// myapp/controller/user_controller.go
package controller

import (
	"net/http"

	"myapp/model"
	"myapp/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.userService.Register(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (uc *UserController) Login(c *gin.Context) {
	var userLogin model.UserLogin
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userService.Login(userLogin)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Set up the session
	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Options(sessions.Options{MaxAge: 60 * 60 * 24}) // Set session to expire after 1 day
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully", "user": user})
}
func (uc *UserController) Enroll(c *gin.Context) {
	learningPathID := c.Param("learningPathID")

	session := sessions.Default(c)
	userID := session.Get("userID")

	err := uc.userService.Enroll(userID.(string), learningPathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrolled in learning path successfully"})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var userUpdate model.UserUpdate
	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	userID := session.Get("userID")

	updatedUser, err := uc.userService.UpdateUser(userID.(string), userUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": updatedUser})
}