package main

import (
	"edu/letu/wan/endpoints"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// read .env file
	godotenv.Load()

	// initialize endpoints
	router := gin.Default() // TODO don't use Default (also add support for .env port)
	endpoints.Initialize(router)
	router.Run()
}