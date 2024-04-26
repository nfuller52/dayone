package environments

import (
	"dayone/application"

	"github.com/gin-gonic/gin"
)

func Development(application *application.ApplicationConfig) {
	// Gin settings
	gin.SetMode(gin.DebugMode)
}
