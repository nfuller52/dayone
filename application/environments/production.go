package environments

import (
	"dayone/application"

	"github.com/gin-gonic/gin"
)

func Production(application *application.ApplicationConfig) {
	// Gin settings
	gin.SetMode(gin.ReleaseMode)
}
