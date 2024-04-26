package main

import (
	"fmt"

	"dayone/api/v1"
	"dayone/application"
	"dayone/application/environments"

	"github.com/gin-gonic/gin"
)

func main() {
	app := application.Start()

	switch app.Env {
	case "production":
		environments.Production(app)
	default:
		environments.Development(app)
	}

	engine := gin.Default()

	v1.Routes(engine)

	serverUrl := fmt.Sprintf("%s:%s", app.Host, app.Port)
	engine.Run(serverUrl)
}
