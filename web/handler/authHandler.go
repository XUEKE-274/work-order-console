package handler

import (
	"github.com/gin-gonic/gin"
	"work-order-console/domain/enum/errorCodeEnum"
	"work-order-console/logger"
	"work-order-console/service"
	"work-order-console/web/response"
)

// 认证
var authHandler = func(logger logger.NewLogger, sessionService service.SessionServiceApi) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log := logger.NewInstance(ctx)
		uri := getUri(ctx)
		log.Info("uri = ", uri)

		// 公共请求不需要认证
		if isPublicPath(ctx) { ctx.Next(); return }
		// 登陆不需要认证
		if isLoginPath(ctx) { ctx.Next(); return }


		token := ctx.GetHeader("Auth")
		// 判断 token 是否存在
		if token == "" {
			log.Info("head miss Auth")
			dealNotLogin(ctx)
			ctx.Abort()
		}else {
			u, success := sessionService.FetchSessionByToken(token)
			if  success {
				// user login success, save context
				ctx.Set("session", u)
				ctx.Next()
			} else {
				log.Info("redis token is empty")
				dealNotLogin(ctx)
				ctx.Abort()
			}
		}
	}
}

func dealNotLogin(ctx *gin.Context) {
	response.ErrorCodeResponse(ctx, errorCodeEnum.NOT_LOGIN)
}