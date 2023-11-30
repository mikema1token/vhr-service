package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vhr-service/service"
)

func main() {
	engine := gin.Default()
	engine.POST("/login", service.UserLogin)
	engine.POST("/logout", service.UserLogout)
	authGroup := engine.Group("/user", func(c *gin.Context) {
		_, err := c.Cookie("token")
		if err != nil {
			c.String(http.StatusUnauthorized, "请先登录")
			return
		} else {
			c.Next()
		}
	})
	authGroup.GET("list-router", service.GetMenuTree)
	authGroup.GET("/list-pos", service.ListPosition)
	authGroup.POST("/add-pos", service.AddPosition)
	engine.Run(":8081")
}
