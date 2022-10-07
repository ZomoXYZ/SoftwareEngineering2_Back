package main

import (
	"edu/letu/wan/endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/authorization", endpoints.Authorization)
	r.POST("/authorization", endpoints.CheckAuthorization)
	
	r.GET("/ping", endpoints.Ping)
        
	r.Run() // listen and serve on 0.0.0.0:8080
}