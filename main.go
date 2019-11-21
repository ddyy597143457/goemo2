package main

import (
	"ddyy/goemo2/app/middleware"

	"ddyy/goemo2/server"

	"github.com/gin-gonic/gin"

	"ddyy/goemo2/router"

)

func main() {
	gin.SetMode("release")
	server.InitLogger()
	server.MysqlInit()
	server.RedisInit()
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(middleware.Logger))
	r.Use(gin.Recovery())
	r = router.SetApiRouter(r)
	r.Run()
}