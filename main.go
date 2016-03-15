package main

import (
	"github.com/ANPez/gogsi/controllers"
	"github.com/ANPez/gogsi/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.Use(middlewares.NewDatabaseMiddleware(os.Getenv("DB_HOST"), os.Getenv("DB_NAME")))
	router.Use(middlewares.NewGoogleMiddleware())
	v1 := router.Group("/v1", middlewares.Auth())
	v1.GET("/test", controllers.Test)

	router.Run(os.Getenv("LISTEN"))
}
