package model

//type BaseModel struct {
//	ID        uint `gorm:"primary_key"`
//	CreatedAt string `gorm:"column:created_at"`
//	UpdatedAt string `gorm:"column:updated_at"`
//	DeletedAt string `gorm:"column:deleted_at"`
//}
type BaseModel struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedAt string `gorm:"column:created_at" json:"created_at" form:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`
	IsDeleted int `gorm:"column:is_deleted" json:"is_deleted" form:"is_deleted"`
}