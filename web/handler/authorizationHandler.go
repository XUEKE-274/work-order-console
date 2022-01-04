package handler

import (
	"github.com/gin-gonic/gin"
	"work-order-console/domain/enum"
	"work-order-console/domain/enum/errorCodeEnum"
	"work-order-console/logger"
	"work-order-console/service"
	"work-order-console/web/response"
)


// 授权
var authorizationHandler = func(logger logger.NewLogger, sessionService service.SessionServiceApi) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log := logger.NewInstance(ctx)
		uri := getUri(ctx)
		log.Info("uri = ", uri)

		// 公共请求不需要认证
		if isPublicPath(ctx) { ctx.Next(); return }

		// 登陆不需要认证
		if isLoginPath(ctx) { ctx.Next(); return }

		// 用户角色
		role := sessionService.ContextSession(ctx).Role

		// 系统管理员 和 操作管理员都是所有权限
		if isAdmin(role) { ctx.Next(); return }

		// 非管理员不能执行用户操作
		if isUserPath(ctx) { response.ErrorCodeResponse(ctx, errorCodeEnum.NOT_PERMISSION); return }

		// 读权限不能执行写操作, 写权限可以执行读权限
		if isReader(role) && isWriterPath(ctx){ response.ErrorCodeResponse(ctx, errorCodeEnum.NOT_PERMISSION); return }

		// 校验业务权限
		if !checkBusinessPermission(ctx, role) { response.ErrorCodeResponse(ctx, errorCodeEnum.NOT_PERMISSION); return }

		// 执行控制器
		ctx.Next()

	}
}

func checkBusinessPermission (ctx *gin.Context, role enum.RoleEnum) bool {
	if isOrderPath(ctx) && isOrderGroup(role) {
		return true
	}
	if isKfPath(ctx) && isKfGroup(role) {
		return true
	}
	if isQtPath(ctx) && isQtGroup(role) {
		return true
	}
	return false
}