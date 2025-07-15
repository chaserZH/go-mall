package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-mall/consts"
	"go-mall/pkg/utils/ctl"
	"go-mall/pkg/utils/log"
	"go-mall/service"
	"go-mall/types"
	"net/http"
)

// UserRegisterHandler 用户注册
func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserRegisterReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		// 参数校验
		if req.Key == "" || len(req.Key) != consts.EncryptMoneyKeyLength {
			err := errors.New("Key长度错误，必须是6位数")
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
		}

		l := service.GetUserSrv()
		resp, err := l.UserRegister(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))

	}
}
