package controller

import (
	"github.com/gin-gonic/gin"
	"ddyy/goemo2/app"
	"ddyy/goemo2/app/model"
)

func GetProductClassifyList(c *gin.Context) {
	var productc model.ProductClassify
	ginres := app.GinResponse{C: c}
	list,err := productc.GetProductClassifyList()
	if err != nil {
		ginres.ResError(app.DATABASE_ERROR,err)
		return
	}
	ginres.ResSuccess(list)
}

func AddProductClassify(c *gin.Context) {
	var productc model.ProductClassify
	ginres := app.GinResponse{C: c}
	if err := c.ShouldBindQuery(&productc); err != nil {
		ginres.ResError(app.QUERY_STRING_ERROR,err)
		return
	}
	err := productc.AddProductClassify()
	if err != nil {
		ginres.ResError(app.DATABASE_ERROR,err)
	}
	ginres.ResSuccess("")
}