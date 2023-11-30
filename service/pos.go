package service

import (
	"github.com/gin-gonic/gin"
	"vhr-service/db"
)

func ListPosition(c *gin.Context) {
	position, err := db.ListPosition()
	if err != nil {
		c.JSON(200, gin.H{"code": "fail", "msg": err.Error()})
	} else {
		c.JSON(200, gin.H{"code": "ok", "data": position})
	}
}

func AddPosition(c *gin.Context) {
	req := struct {
		Name string
	}{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, gin.H{"code": "fail", "msg": err.Error()})
	} else {
		err = db.AddPosition(req.Name)
		if err != nil {
			c.JSON(200, gin.H{"code": "fail", "msg": err.Error()})
		} else {
			c.JSON(200, gin.H{"code": "ok", "data": "ok"})
		}
	}
}
