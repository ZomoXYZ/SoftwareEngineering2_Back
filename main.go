package main

import (
	"edu/letu/wan/endpoints"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// read .env file
	godotenv.Load()

	// initialize REST endpoints
	r := gin.Default()
	endpoints.Initialize(r)
	r.Run()
}