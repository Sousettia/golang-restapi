package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sousettia/go-restapi/config"
	"github.com/sousettia/go-restapi/handlers"
	"github.com/sousettia/go-restapi/middlewares"
	"github.com/sousettia/go-restapi/tasks"
)

func main() {

	//Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//Connect to MongoDB
	config.ConnectDB()

	//Connect to Background Task
	go tasks.StartUserCountLogger()
	//Create Gin router
	r := gin.Default()

	//Connect Logger
	r.Use(middlewares.Logger())

	//Routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	//Protected routes
	api := r.Group("/api")
	api.Use(middlewares.JwtVerify())
	{
		api.GET("/users", handlers.GetUsers)
		api.GET("/users/:id", handlers.GetUserByID)
		api.PUT("/users/:id", handlers.UpdateUser)
		api.DELETE("/users/:id", handlers.DeleteUser)
	}

	log.Println("🚀 Server running at http://localhost:8080")
	r.Run(":8080")
}
