package dao

import (
	"context"
	"go-mall/repository/db/model"
	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

// GetOrderById gorm dao层查询 1、链式调用
func (dao *OrderDao) GetOrderById(id, uid uint) (r *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id = ? and user_id = ?", id, uid).First(&r).Error
	return
}

// CreateOrder 创建订单
func (dao *OrderDao) CreateOrder(order *model.Order) (err error) {
	return dao.DB.Create(&order).Error
}
