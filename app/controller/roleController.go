package controller

import (
	"github.com/gin-gonic/gin"
	"ddyy/goemo2/app"
	"ddyy/goemo2/app/model"
	"strconv"
)

func AddRole(c *gin.Context) {
	var role model.Role
	ginres := app.GinResponse{C: c}
	if err := c.ShouldBindQuery(&role); err != nil {
		ginres.ResError(app.QUERY_STRING_ERROR,err)
		return
	}
	err := role.AddRole()
	if err != nil {
		ginres.ResError(app.DATABASE_ERROR,err)
		return
	}
	ginres.ResSuccess("添加角色成功")
}

func DelRole(c *gin.Context) {
	var role model.Role
	ginres := app.GinResponse{C: c}
	roleId := c.DefaultQuery("id","")
	if roleId == "" {
		ginres.ResError(app.QUERY_STRING_ERROR,nil)
	}
	roleIdint,_:= strconv.Atoi(roleId)
	err := role.DelRole(roleIdint)
	if err != nil {
		ginres.ResError(app.DATABASE_ERROR,err)
		return
	}
	ginres.ResSuccess("删除角色成功")
}