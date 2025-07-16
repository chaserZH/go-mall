package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/pkg/utils/log"
	"go-mall/types"
	"net/http"
)

func CreateOrderHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.OrderCreateReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		//l := service.GetOrderSrv()
		//resp, err := l.OrderCreate(ctx.Request.Context(), &req)

	}
}
