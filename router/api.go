package router

import (
	"github.com/gin-gonic/gin"
	"ddyy/goemo2/app/controller"
	"ddyy/goemo2/app/middleware"
)

func SetApiRouter(r *gin.Engine) *gin.Engine {
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome")
	})
	r.POST("/user/register", controller.Register)
	r.POST("/user/login", controller.Login)
	r.POST("/user/loginout", controller.LoginOut)

	role := r.Group("/role")
	role.Use(middleware.Loginer(),middleware.Roler())
	{
		role.POST("/add", controller.AddRole)
		role.POST("/del", controller.DelRole)
	}

	r.GET("/productclassify/list", controller.GetProductClassifyList)
	r.POST("/productclassify/add", controller.AddProductClassify)

	r.GET("/product/list", controller.GetProductList)
	r.POST("/product/add", controller.AddProduct)
	r.GET("/product/info", controller.GetProductInfo)
	r.GET("/product/del", controller.DelProduct)
	return r
}
