package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourproject/models"
	"github.com/yourproject/services"
)

type InstructorController interface {
	RegisterInstructor(ctx *gin.Context)
	LoginInstructor(ctx *gin.Context)
	UpdateInstructor(ctx *gin.Context)
	ForgotPassword(ctx *gin.Context)
	CreateLearningPath(ctx *gin.Context)
	UpdateLearningPath(ctx *gin.Context)
	DeleteLearningPath(ctx *gin.Context)
	CreateExercise(ctx *gin.Context)
	UpdateExercise(ctx *gin.Context)
	DeleteExercise(ctx *gin.Context)
	GetEnrolledStudents(ctx *gin.Context)
}

type instructorController struct {
	instructorService services.InstructorService
	authService       services.AuthService
}

func NewInstructorController(is services.InstructorService, as services.AuthService) InstructorController {
	return &instructorController{
		instructorService: is,
		authService:       as,
	}
}
// RegisterInstructor ...
func (c *instructorController) RegisterInstructor(ctx *gin.Context) {
	var instructor models.Instructor
	if err := ctx.ShouldBindJSON(&instructor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := services.HashPassword(instructor.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	instructor.Password = hashedPassword
	result, err := c.instructorService.CreateInstructor(instructor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// LoginInstructor ...
func (c *instructorController) LoginInstructor(ctx *gin.Context) {
	var credentials models.Credentials
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	instructor, err := c.instructorService.GetInstructorByEmail(credentials.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !services.CheckPasswordHash(credentials.Password, instructor.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := c.authService.GenerateToken(instructor.ID.Hex(), "instructor")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// UpdateInstructor ...
func (c *instructorController) UpdateInstructor(ctx *gin.Context) {
	var instructor models.Instructor
	if err := ctx.ShouldBindJSON(&instructor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	instructorID, err := c.authService.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied"})
		return
	}

	instructor.ID = instructorID
	result, err := c.instructorService.UpdateInstructor(instructor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
// ForgotPassword ...
func (c *instructorController) ForgotPassword(ctx *gin.Context) {
	var emailRequest models.EmailRequest
	if err := ctx.ShouldBindJSON(&emailRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.instructorService.SendForgotPasswordEmail(emailRequest.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password reset email sent"})
}

// CreateLearningPath ...
func (c *instructorController) CreateLearningPath(ctx *gin.Context) {
	var learningPath models.LearningPath
	if err := ctx.ShouldBindJSON(&learningPath); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	instructorID, err := c.authService.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied"})
		return
	}

	result, err := c.instructorService.CreateLearningPath(instructorID, learningPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// UpdateLearningPath ...
func (c *instructorController) UpdateLearningPath(ctx *gin.Context) {
	var learningPath models.LearningPath
	if err := ctx.ShouldBindJSON(&learningPath); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	instructorID, err := c.authService.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied"})
		return
	}

	err = c.instructorService.UpdateLearningPath(instructorID, learningPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Learning path updated"})
}

// DeleteLearningPath ...
func (c *instructorController) DeleteLearningPath(ctx *gin.Context) {
	learningPathID := ctx.Param("id")

	instructorID, err := c.authService.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied"})
		return
	}

	err = c.instructorService.DeleteLearningPath(instructorID, learningPathID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Learning path deleted"})
}

// CreateExercise ...
func (c *instructorController) CreateExercise(ctx *gin.Context) {
	var exercise models.Exercise
	if err := ctx.ShouldBindJSON(&exercise); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	instructorID, err := c.authService.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied"})
		return
	}

	learningPathID := ctx.Param("path_id")

	result, err := c.instructorService.CreateExercise(instructorID, learningPathID, exercise)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
// UpdateExercise ...
func (c *instructorController) UpdateExercise(ctx *gin.Context) {
	var exercise models.Exercise
	if err := ctx.ShouldBindJSON(&exercise); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	instructorID, err := c.authService.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied"})
		return
	}

	learningPathID := ctx.Param("path_id")

	err = c.instructorService.UpdateExercise(instructorID, learningPathID, exercise)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Exercise updated"})
}

// DeleteExercise ...
func (c *instructorController) DeleteExercise(ctx *gin.Context) {
	exerciseID := ctx.Param("id")

	instructorID, err := c.authService.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied"})
		return
	}

	learningPathID := ctx.Param("path_id")

	err = c.instructorService.DeleteExercise(instructorID, learningPathID, exerciseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Exercise deleted"})
}

// GetEnrolledStudents ...
func (c *instructorController) GetEnrolledStudents(ctx *gin.Context) {
	learningPathID := ctx.Param("path_id")

	instructorID, err := c.authService.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied"})
		return
	}

	students, err := c.instructorService.GetEnrolledStudents(instructorID, learningPathID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, students)
}