package controller

import (
	"github.com/gin-gonic/gin"
	"ddyy/goemo2/app"
	"ddyy/goemo2/app/model"
	"strconv"
)

func GetProductList(c *gin.Context) {
	var product model.Product
	ginres := app.GinResponse{C: c}
	classifyid := c.DefaultQuery("classify_id","")
	var classifyidint int
	if classifyid == "" {
		classifyidint = 0
	} else {
		classifyidint,_ = strconv.Atoi(classifyid)
	}
	list,err := product.GetProductList(classifyidint)
	if err != nil {
		ginres.ResError(app.DATABASE_ERROR,err)
	}
	ginres.ResSuccess(list)
}

func AddProduct(c *gin.Context) {
	var product model.Product
	ginres := app.GinResponse{C: c}
	if err := c.ShouldBindQuery(&product); err != nil {
		ginres.ResError(app.QUERY_STRING_ERROR,err)
		return
	}
	err := product.AddProduct()
	if err != nil {
		ginres.ResError(app.DATABASE_ERROR,err)
	}
	ginres.ResSuccess("")
}

func GetProductInfo(c *gin.Context) {
	var product model.Product
	ginres := app.GinResponse{C: c}
	id := c.DefaultQuery("id","")
	if id == "" {
		ginres.ResError(app.QUERY_STRING_ERROR,nil)
		return
	}
	idint,_:= strconv.Atoi(id)
	res ,err := product.GetProductInfo(idint)
	if err != nil {
		ginres.ResError(app.DATA_NOT_FOUND,nil)
		return
	}
	ginres.ResSuccess(res)
}

func DelProduct(c *gin.Context) {
	var product model.Product
	ginres := app.GinResponse{C: c}
	id := c.DefaultQuery("id","")
	if id == "" {
		ginres.ResError(app.QUERY_STRING_ERROR,nil)
		return
	}
	idint,_:= strconv.Atoi(id)
	err := product.DelProduct(idint)
	if err != nil {
		ginres.ResError(app.DATA_NOT_FOUND,err)
		return
	}
	ginres.ResSuccess("")
}
