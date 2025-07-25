package model

import "gorm.io/gorm"

// 购物车模型
type Cart struct {
	gorm.Model
	UserId    uint
	ProductId uint `gorm:"not null"`
	BossID    uint
	Num       uint
	MaxNum    uint
	Check     bool
}
