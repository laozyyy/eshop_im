package model

import "time"

type GoodsSku struct {
	ID         int64     `gorm:"primaryKey;column:id"`
	Sku        string    `gorm:"column:sku"`
	GoodsID    string    `gorm:"column:goods_id"`
	TagID      string    `gorm:"column:tag_id"`
	Name       string    `gorm:"column:name"`
	Price      float64   `gorm:"column:price"`
	Spec       string    `gorm:"column:spec"`
	ShowPic    string    `gorm:"column:show_pic"`
	DetailPic  string    `gorm:"column:detail_pic"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
	IsDeleted  int32     `gorm:"column:is_deleted"`
}
