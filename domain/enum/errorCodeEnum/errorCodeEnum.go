package errorCodeEnum

import "work-order-console/domain/enum"

const (
	TEST_ERROR enum.ErrorCodeEnum = "TEST_ERROR"

	PARAMS_ERROR           enum.ErrorCodeEnum = "PARAMS_ERROR"
	NOT_LOGIN              enum.ErrorCodeEnum = "NOT_LOGIN"
	AUTH_FAIL              enum.ErrorCodeEnum = "AUTH_FAIL"
	NOT_PERMISSION         enum.ErrorCodeEnum = "NOT_PERMISSION"
	SYS_USER_CAN_NOT_OPEAT enum.ErrorCodeEnum = "SYS_USER_CAN_NOT_OPEAT"
	USERNAME_HAS_EXIST     enum.ErrorCodeEnum = "USERNAME_HAS_EXIST"
	USERNAME_NOT_EXIST     enum.ErrorCodeEnum = "USERNAME_NOT_EXIST"
)

var mapping = map[enum.ErrorCodeEnum]string{
	TEST_ERROR: "测试错误",

	PARAMS_ERROR:           "参数错误",
	NOT_LOGIN:              "未登陆",
	AUTH_FAIL:              "认证失败",
	NOT_PERMISSION:         "没有权限",
	SYS_USER_CAN_NOT_OPEAT: "系统管理员不能操作",
	USERNAME_HAS_EXIST:     "用户已经存在",
	USERNAME_NOT_EXIST:     "用户名不存在",
}

func GetValueByCode(code enum.ErrorCodeEnum) string {
	return mapping[code]
}
