package model

import (
	"encoding/json"
	"github.com/pkg/errors"
	"ddyy/goemo2/helper"
	"ddyy/goemo2/server"
	"time"
)

type User struct {
	BaseModel
	Name    string `gorm:"size:20" form:"name"`
	Phone   string `gorm:"size:15" form:"phone" binding:"required"`
	Passwd  string `gorm:"size:50" form:"passwd" binding:"required"`
	Salt    string `gorm:"size:15"`
}

func (u *User) Register() error {
	var user User
	db := server.GetDBEngine()
	user.Name = u.Name
	user.Phone = u.Phone
	user.Salt = helper.RandString(10)
	user.Passwd = helper.Sha1(u.Passwd)
	createdAt := helper.TimeFormat(time.Now())
	user.CreatedAt = createdAt
	user.UpdatedAt = createdAt
	db.Create(&user)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (u *User)Login() (string,error) {
	var user User
	db := server.GetDBEngine()
	phone := u.Phone
	passwd := helper.Sha1(u.Passwd)
	db.Where("phone=? and passwd=?",phone,passwd).Select("id,phone,name,salt").Find(&user)
	if db.Error != nil {
		return "",db.Error
	}
	token := helper.Sha1(user.Phone+user.Salt)
	userInfobytes ,err := json.Marshal(user)
	if err != nil {
		return "",err
	}
	redisConn := server.GetRedisConn()
	ex, _ := redisConn.Do("SET", "user_"+token,string(userInfobytes),"EX",60*60*24)
	if ex == nil {
		return "",errors.New("")
	}
	return token,nil
}

func LoginOut(token string) error {
	redisConn := server.GetRedisConn()
	ex, _ := redisConn.Do("DEL", "user_"+token)
	if ex == nil {
		return errors.New("")
	}
	return nil
}