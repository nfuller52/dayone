package main

import (
	"dayone/application"
	"dayone/application/environments"
	v1 "dayone/src/api/v1"

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

	engine.Run(app.ServerUrl())
}
