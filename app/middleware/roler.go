package middleware

import (
	"encoding/json"
	"errors"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"ddyy/goemo2/app"
	"ddyy/goemo2/app/model"
	"ddyy/goemo2/server"
)

func Roler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.DefaultQuery("token", "")
		ginres := app.GinResponse{C: c}
		redisConn := server.GetRedisConn()
		v,err := redis.String(redisConn.Do("GET", "user_"+token))
		if err != nil {
			ginres.ResError(app.DATABASE_ERROR,errors.New("redis异常"))
			c.Abort()
			return
		}
		var userinfo model.UserLoginInfo
		err = json.Unmarshal([]byte(v),&userinfo)
		if err != nil {
			ginres.ResError(app.DATABASE_ERROR,errors.New("loginuserinfo解析异常"))
			c.Abort()
			return
		}
		if userinfo.RoleId != app.SYSTEM_MANAGER {
			ginres.ResError(app.AUTH_ERROR,nil)
			c.Abort()
			return
		}
	}
}