package middleware

import (
	"errors"
	"ddyy/goemo2/app"
	"ddyy/goemo2/server"

	"github.com/gin-gonic/gin"
)

func Loginer() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.DefaultQuery("token", "")
		ginres := app.GinResponse{C: c}
		if token == "" {
			ginres.ResError(app.MISS_TOKEN,errors.New("token缺失"))
			c.Abort()
			return
		}
		redisConn := server.GetRedisConn()
		ex, err := redisConn.Do("GET", "user_"+token)
		if err != nil {
			ginres.ResError(app.DATABASE_ERROR,err)
			c.Abort()
			return
		}
		if ex == nil {
			ginres.ResError(app.USER_NOT_LOGIN,nil)
			c.Abort()
			return
		}
	}
}
