package main

import (
	"fmt"
	"log"
	"app1_be/config"
	"app1_be/controller"
	"app1_be/repository"
	"app1_be/router"
	"app1_be/service"
	"net/http"
)

func main() {
	mongoDBConfig := config.GetMongoDBConfig()
	mongoClient, err := database.ConnectMongoDB(mongoDBConfig)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer mongoClient.Disconnect(nil)

	// Set up repositories, services, and controllers using the connected MongoDB client
	userRepo := repository.NewUserRepository(mongoClient, mongoDBConfig.Database)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// ... Set up other repositories, services, and controllers for different entities

	// Create and configure the router
	appRouter := router.NewRouter(userController) // Add other controllers as needed

	// Start the server
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", appRouter))
}