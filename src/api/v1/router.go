package v1

import (
	"dayone/src/api/v1/auth"

	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine) {
	router := engine.Group("/api/v1")

	auth.AuthRoutes(router)
}
