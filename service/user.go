package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"vhr-service/db"
)

func UserLogin(ctx *gin.Context) {
	q := struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}{}
	err := ctx.ShouldBind(&q)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	userModel, err := db.GetUserModelByName(q.UserName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	if userModel.Password == q.Password {
		ctx.SetCookie("token", q.UserName+q.Password, 3600*24, "/", "localhost", false, true)
		ctx.JSON(http.StatusOK, userModel)
		return
	} else {
		ctx.JSON(http.StatusInternalServerError, errors.New("invalid username or password"))
		return
	}
}

func UserLogout(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(500, gin.H{"err": err})
		return
	}
	c.SetCookie("token", cookie, 1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"status": "ok"})
}
