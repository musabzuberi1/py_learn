package controller

import (
	"myapp/model"
	"myapp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	service *service.AdminService
}

func (ac *AdminController) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := ac.service.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ac *AdminController) GetInstructor(c *gin.Context) {
	id := c.Param("id")
	instructor, err := ac.service.GetInstructor(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, instructor)
}

func (ac *AdminController) GetAllUsers(c *gin.Context) {
	users, err := ac.service.GetAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (ac *AdminController) GetAllInstructors(c *gin.Context) {
	instructors, err := ac.service.GetAllInstructors(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, instructors)
}

func (ac *AdminController) UpdateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	if err := ac.service.UpdateUser(c.Request.Context(), id, user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (ac *AdminController) UpdateInstructor(c *gin.Context) {
	var instructor model.Instructor
	if err := c.ShouldBindJSON(&instructor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	if err := ac.service.UpdateInstructor(c.Request.Context(), id, instructor); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (ac *AdminController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := ac.service.DeleteUser(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func (ac *AdminController) DeleteInstructor(c *gin.Context) {
	id := c.Param("id")
	if err := ac.service.DeleteInstructor(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
