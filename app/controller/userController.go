package controller

import (
	"ddyy/goemo2/app"
	"ddyy/goemo2/app/model"

	"github.com/gin-gonic/gin"

	"errors"
)

func Register(c *gin.Context) {
	var user model.User
	ginres := app.GinResponse{C: c}
	if err := c.ShouldBindQuery(&user); err != nil {
		ginres.ResError(app.QUERY_STRING_ERROR,err)
		return
	}
	err := user.Register()
	if err != nil {
		ginres.ResError(app.DATABASE_ERROR,err)
		return
	}
	ginres.ResSuccess("")
}

func Login(c *gin.Context) {
	var user model.User
	ginres := app.GinResponse{C: c}
	if err := c.ShouldBindQuery(&user); err != nil {
		ginres.ResError(app.QUERY_STRING_ERROR,err)
		return
	}
	token,err := user.Login()
	if err != nil {
		ginres.ResError(app.DATABASE_ERROR,err)
		return
	}
	ginres.ResSuccess(token)
}

func LoginOut(c *gin.Context) {
	ginres := app.GinResponse{C: c}
	token := c.DefaultQuery("token","")
	if token == "" {
		ginres.ResError(app.MISS_TOKEN,errors.New("token缺失"))
		return
	}
	err := model.LoginOut(token)
	if err != nil {
		ginres.ResError(app.DATABASE_ERROR,err)
	}
	ginres.ResSuccess("退出成功")
}