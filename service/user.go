package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"vhr-service/db"
)

func UserLogin(ctx *gin.Context) {
	q := struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}{}
	err := ctx.ShouldBind(&q)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	userModel, err := db.GetUserModelByName(q.UserName)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	if userModel.Password == q.Password {
		ctx.JSON(200, gin.H{"message": "ok"})
		return
	} else {
		ctx.JSON(500, errors.New("password incorrect"))
		return
	}
}
