package model

import (
	"ddyy/goemo2/helper"
	"ddyy/goemo2/server"
	"time"
)

const (
	NOT_DELETED = 0
	DELETED = 1
)

type Product struct {
	BaseModel
	ProductName    string `gorm:"size:20" json:"product_name" form:"product_name"`
	ClassifyId int `json:"classify_id" form:"classify_id"`
	Type string `gorm:"size:20" json:"type" form:"type"`
}

func (p Product)GetAllProductList() ([]map[string]string,error){
	var list []map[string]string
	db := server.GetDBEngine()
	db.Table("product_classifies")
	rows, err := db.Table("product_classifies").Select("product_classifies.id,product_classifies.classify_name,product_classifies.pid,products.product_name").Joins("left join products on product_classifies.id = products.classify_id").Rows()
	if err != nil {
		return list,err
	}
	var s [4]string
	var m map[string]string
	for rows.Next() {
		rows.Scan(&s[0],&s[1],&s[2],&s[3])
		m = map[string]string{
			"id":s[0],
			"classify_name":s[1],
			"pid":s[2],
			"product_name":s[3],
		}
		list = append(list, m)
	}
	return list,nil
}

func (p Product)GetProductList(classifyid int) ([]Product,error){
	var pm []Product
	db := server.GetDBEngine()
	if classifyid <= 0 {
		if err := db.Find(&pm).Error; err != nil {
			return pm,err
		}
	} else {
		if err := db.Where("classify_id=?",classifyid).Find(&pm).Error; err != nil {
			return pm,err
		}
	}
	return pm,nil
}

func (p Product)AddProduct() error{
	var pm Product
	pm.ProductName = p.ProductName
	pm.ClassifyId = p.ClassifyId
	pm.Type = p.Type
	createdAt := helper.TimeFormat(time.Now())
	pm.CreatedAt = createdAt
	pm.UpdatedAt = createdAt
	db := server.GetDBEngine()
	if err := db.Create(&pm).Error; err != nil {
		return err
	}
	return nil
}

func (Product)GetProductInfo(id int) (Product,error) {
	var pm Product
	db := server.GetDBEngine()
	if err := db.Where("id=? and is_deleted=?",id,NOT_DELETED).First(&pm).Error;err != nil {
		return pm,err
	}
	return pm,nil
}

func (Product)DelProduct(id int) (error) {
	db := server.GetDBEngine()
	if err := db.Table("products").Where("id=?",id).Update("is_deleted",DELETED).Error;err != nil {
		return err
	}
	return nil
}