package model

import (
	"ddyy/goemo2/helper"
	"ddyy/goemo2/server"
	"time"
)

type ProductClassify struct {
	BaseModel
	ClassifyName    string `gorm:"size:20" json:"classify_name" form:"classify_name"`
	Pid   int `json:"pid" form:"pid"`
	Children []ProductClassify
	Item []Product
}

func (ProductClassify) GetProductClassifyList() (interface{},error){
	var pcms []ProductClassify
	db := server.GetDBEngine()
	if err := db.Find(&pcms).Error; err != nil {
		return pcms,err
	}
	if len(pcms) > 0 {
		product := Product{}
		productList,_ := product.GetProductList(0)
		for id,v := range pcms {
			for _,v1 := range productList {
				if v1.ClassifyId == v.ID {
					pcms[id].Item = append(pcms[id].Item,v1)
				}
			}
		}
		data := buildTreeData(pcms)
		pcms = makeTreeCore(0, data)
	}
	return pcms,nil
}

func buildTreeData(data []ProductClassify) map[int]map[int]ProductClassify{
	makedata := make(map[int]map[int]ProductClassify)
	for _,v := range data {
		id := v.ID
		pid := v.Pid
		if _,ok := makedata[pid]; !ok {
			makedata[pid] = make(map[int]ProductClassify)
		}
		makedata[pid][id] = v
	}
	return makedata
}

func makeTreeCore(index int,data map[int]map[int]ProductClassify) []ProductClassify {
	tmp := make([]ProductClassify,0)
	for id, item := range data[index] {
		if data[id] != nil {
			item.Children = makeTreeCore(id, data)
		}
		tmp = append(tmp, item)
	}
	return tmp
}

func (ProductClassify)makeProductClassifyTree() {
	
}

func (pc ProductClassify)AddProductClassify() error{
	var pcm ProductClassify
	pcm.ClassifyName = pc.ClassifyName
	pcm.Pid = pc.Pid
	createdAt := helper.TimeFormat(time.Now())
	pcm.CreatedAt = createdAt
	pcm.UpdatedAt = createdAt
	db := server.GetDBEngine()
	if err := db.Create(&pcm).Error; err != nil {
		return err
	}
	return nil
}