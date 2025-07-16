package model

import "gorm.io/gorm"

// Order 订单信息
type Order struct {
	gorm.Model
	UserID    uint   `gorm:"not null"`
	ProductID uint   `gorm:"not null"`
	BossID    uint   `gorm:"not null"`
	AddressID uint   `gorm:"not null"`
	Num       int    // 数量
	OrderNum  uint64 // 订单号
	Type      uint   // 1 未支付  2 已支付
	Money     float64
}

// TableName 自定义表名
func (Order) TableName() string {
	return "t_order" // 指定表名为 order 而不是默认的 orders
}
