package service

import (
	"sync"
)

var OrderSrvIns *OrderSrv
var OrderSrvOnce sync.Once

type OrderSrv struct {
}

func GetOrderSrv() *OrderSrv {
	OrderSrvOnce.Do(func() {
		OrderSrvIns = &OrderSrv{}
	})
	return OrderSrvIns
}

//func (s *OrderSrv) OrderCreate(ctx *gin.Context, req *types.OrderCreateReq) (resp interface{}, err error) {
//	u, err := ctl.Response(ctx)
//
//}
