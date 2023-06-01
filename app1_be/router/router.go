// myapp/router/router.go
package router

import (
	"myapp/controller"
	"myapp/middleware"
	"myapp/mongostore"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func NewRouter(userController *controller.UserController, userProgressController *controller.UserProgressController, learningPathController *controller.LearningPathController) *gin.Engine {
	router := gin.Default()

	// Set up session handling using MongoDB
	store, err := mongostore.NewMongoStore("mongodb://localhost:27017", "your_db_name", "sessions", 60*60*24, true, []byte("your_session_secret"))
	if err != nil {
		panic(err)
	}
	router.Use(sessions.Sessions("mysession", store))

	// User routes
	router.POST("/api/users/register", userController.Register)
	router.POST("/api/users/login", userController.Login)
	router.POST("/api/users/forgot-password", userController.ForgotPassword)


	// Instructor routes
	router.POST("/api/instructors/register", instructorController.Register)
	router.POST("/api/instructors/login", instructorController.Login)
	router.POST("/api/instructors/forgot-password", instructorController.ForgotPassword)



	// Routes that require authentication
	authRoutes := router.Group("/")
	authRoutes.Use(middleware.Auth())
	{
		authRoutes.POST("/api/enroll/:learningPathID", userController.Enroll)
		authRoutes.PUT("/api/user", userController.UpdateUser)
		authRoutes.POST("/api/user-progress/:exerciseID", userProgressController.SaveUserCode)
		authRoutes.GET("/api/user/points", userController.GetPointsAndLevel)
		authRoutes.POST("/api/logout", userController.Logout)
	}

	instructorAuthRoutes := router.Group("/")
	instructorAuthRoutes.Use(middleware.InstructorAuth())
	{
		instructorAuthRoutes.PUT("/api/instructors/update-email", instructorController.UpdateEmail)
		instructorAuthRoutes.PUT("/api/instructors/change-password", instructorController.ChangePassword)

		// Learning path CRUD
		instructorAuthRoutes.POST("/api/learning-paths", learningPathController.CreateLearningPath)
		instructorAuthRoutes.GET("/api/learning-paths", learningPathController.GetAllLearningPaths)
		instructorAuthRoutes.GET("/api/learning-paths/:pathID", learningPathController.GetLearningPath)
		instructorAuthRoutes.PUT("/api/learning-paths/:pathID", learningPathController.UpdateLearningPath)
		instructorAuthRoutes.DELETE("/api/learning-paths/:pathID", learningPathController.DeleteLearningPath)

		// Exercise CRUD within learning paths
		instructorAuthRoutes.POST("/api/learning-paths/:pathID/exercises", learningPathController.CreateExercise)
		instructorAuthRoutes.GET("/api/learning-paths/:pathID/exercises", learningPathController.GetAllExercises)
		instructorAuthRoutes.GET("/api/learning-paths/:pathID/exercises/:exerciseID", learningPathController.GetExercise)
		instructorAuthRoutes.PUT("/api/learning-paths/:pathID/exercises/:exerciseID", learningPathController.UpdateExercise)
		instructorAuthRoutes.DELETE("/api/learning-paths/:pathID/exercises/:exerciseID", learningPathController.DeleteExercise)

		// View enrolled students
		instructorAuthRoutes.GET("/api/learning-paths/:pathID/students", learningPathController.GetEnrolledStudents)
	}

	return router
}
