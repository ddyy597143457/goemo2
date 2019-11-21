package model

import (
	"ddyy/goemo2/helper"
	"ddyy/goemo2/server"
	"time"
)

type Role struct {
	BaseModel
	RoleName    string `gorm:"size:20" form:"role_name"`
}

func (r Role)AddRole() error {
	var role Role
	role.RoleName = r.RoleName
	createdAt := helper.TimeFormat(time.Now())
	role.CreatedAt = createdAt
	role.UpdatedAt = createdAt
	db := server.GetDBEngine()
	if err := db.Create(&role).Error;err != nil {
		return err
	}
	return nil
}

func (Role)DelRole(roleid int) error {
	db := server.GetDBEngine()
	if err := db.Table("roles").Where("id=?",roleid).Update("is_deleted",DELETED).Error;err != nil {
		return err
	}
	return nil
}