package handler

import (
	"github.com/gin-gonic/gin"
	"strings"
	"work-order-console/domain/enum"
	"work-order-console/domain/enum/roleEnum"
	"work-order-console/web/router"
)

func getUri(ctx *gin.Context) string {
	return ctx.Request.RequestURI
}

// isLoginPath
// 是否是登陆接口
func isLoginPath(ctx *gin.Context) bool {
	uri := getUri(ctx)
	return strings.HasPrefix(uri, router.Login)

}

// isPublicPath
// 是否为公共接口， 不需要认证
// SwaggerPath Healthy System
func isPublicPath(ctx *gin.Context) bool {
	uri := getUri(ctx)
	return strings.HasPrefix(uri, router.SwaggerPath) || strings.HasPrefix(uri, router.Healthy) || strings.HasPrefix(uri, router.System)
}

// isReaderPath
// 是否是读接口
func isReaderPath(ctx *gin.Context) bool {
	return ctx.Request.Method == "GET"
}

// isWriterPath
// 是否是写接口
func isWriterPath(ctx *gin.Context) bool  {
	return ctx.Request.Method != "GET"
}

// isUserPath
// 是否是用户请求
func isUserPath(ctx *gin.Context) bool {
	uri := getUri(ctx)
	return strings.HasPrefix(uri, router.User)
}

func isOrderPath(ctx *gin.Context) bool {
	uri := getUri(ctx)
	return strings.HasPrefix(uri, router.ORDER)
}

func isKfPath(ctx *gin.Context) bool {
	uri := getUri(ctx)
	return strings.HasPrefix(uri, router.KF)
}

func isQtPath(ctx *gin.Context) bool {
	uri := getUri(ctx)
	return strings.HasPrefix(uri, router.QT)
}




// 角色

func getOperation(role enum.RoleEnum) string {
	str := string(role)
	return strings.Split(str, "_")[0]
}

func getGroup(role enum.RoleEnum) string {
	str := string(role)
	return strings.Split(str, "_")[1]
}


// isAdmin
// 是否是超级管理员
func isAdmin(role enum.RoleEnum) bool {
	return roleEnum.SYSTEM == role || roleEnum.ADMIN == role
}
// isReader
func isReader(role enum.RoleEnum) bool {
	return "READER" == getOperation(role)
}

// isWriter
func isWriter(role enum.RoleEnum) bool {
	return "WRITER" == getOperation(role)
}

// isOrderGroup
func isOrderGroup(role enum.RoleEnum) bool {
	return "ORDER" == getGroup(role)
}

// isKfGroup
func isKfGroup(role enum.RoleEnum) bool {
	return "KF" == getGroup(role)
}

// isQtGroup
func isQtGroup(role enum.RoleEnum) bool {
	return "QT" == getGroup(role)
}


