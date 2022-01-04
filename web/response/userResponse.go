package response

import "work-order-console/domain/enum"

type UserDetailResponse struct {
	Id string `json:"id"`  // 用户id
	Username string `json:"username"` // 用户名
	Role enum.RoleEnum `json:"role"` // 用户角色

}


type UserListResponse struct {
	Users *[]UserDetailResponse `json:"users"` // 用户数组
	Total int64 `json:"total"`  // 总数
}
