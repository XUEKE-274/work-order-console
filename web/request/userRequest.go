package request

import "work-order-console/domain/enum"


type UserAddRequest struct {
	Username string `json:"username" valid:"required"` // 用户名
	Password string `json:"password" valid:"required"` // 密码
	Role enum.RoleEnum `json:"role" valid:"required"` // 角色
}

type UserModifyRequest struct {
	Role enum.RoleEnum `json:"role" valid:"required"` // 角色
}

type UserListRequest struct {
	PageRequest
	Username string `json:"username"` // 用户名称
	Role enum.RoleEnum `json:"role"`  // 角色
}
