package main

import (
	"github.com/gin-gonic/gin"
	"vhr-service/service"
)

func main() {
	engine := gin.Default()
	registerRouter(engine)
	engine.Run()
}

func registerRouter(engine *gin.Engine) {
	engine.POST("login", service.UserLogin)
}
