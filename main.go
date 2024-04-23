package main

import (
	"fmt"
	"net/http"

	"dayone/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.LoadApplicationConfig()

	router := gin.Default()

	if config.Env == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	serverPort := fmt.Sprintf(":%s", config.Port)

	router.Run(serverPort)
}
