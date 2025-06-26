package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	setupRouters(router)
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(router.Run(":" + port)) //fatal is used to To log the error and stop the app if starting the server fails and we use ":" to with out : This tries to bind to a UNIX socket named "8080", not a TCP port! That's not what you want.:8080 tells Gin to listen on TCP port 8080 on all interfaces (i.e., 0.0.0.0:8080).This is the correct way to start an HTTP server on a specific port.

}

func setupRouters(router *gin.Engine) {
	router.POST("/api/v1/shorten", routes.shorten)
}
