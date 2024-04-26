package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine) {
	router := engine.Group("/api/v1")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
