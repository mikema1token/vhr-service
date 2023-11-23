package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vhr-service/service"
)

func main() {
	engine := gin.Default()
	engine.POST("/login", service.UserLogin)
	authGroup := engine.Group("/user", func(c *gin.Context) {
		_, err := c.Cookie("token")
		if err != nil {
			c.String(http.StatusUnauthorized, "请先登录")
			return
		} else {
			c.Next()
		}
	})
	authGroup.POST("", nil)
	engine.Run(":8081")
}

func registerAuthApiRouter(engine *gin.Engine) {

}
