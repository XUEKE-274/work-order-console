package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"strconv"
	"work-order-console/domain/enum"
	"work-order-console/logger"
	"work-order-console/service"
	"work-order-console/web/request"
	"work-order-console/web/response"
)

type UserControllerApi interface {
	AddUser(c *gin.Context)
	ModifyUser(c *gin.Context)
	DelUser(c *gin.Context)
	ListUser(c *gin.Context)
	DetailUser(c *gin.Context)
}
type userController struct {
	logger      logger.NewLogger
	userService service.UserServiceApi
}

var regUserController = fx.Provide(func(logger logger.NewLogger, userService service.UserServiceApi) UserControllerApi {
	return &userController{
		logger,
		userService,
	}
})

// AddUser
// @Tags 用户管理
// @Summary 创建用户
// @Description 创建用户
// @Accept  json
// @Produce  json
// @Param param body request.UserAddRequest true "body"
// @Param Auth header string true "Token"
// @Success 200 {object} response.Response "desc"
// @Router /user [post]
func (mine *userController) AddUser(c *gin.Context) {

	log := mine.logger.NewInstance(c)

	params := &request.UserAddRequest{}
	e := c.BindJSON(&params)
	if e != nil {
		log.Error("json parse error")
		response.ParamsErrorResponse(c)
		return
	}
	success, _ := govalidator.ValidateStruct(params)
	if !success {
		log.Error("params is empty = ", params)
		response.ParamsErrorResponse(c)
		return
	}
	e = mine.userService.AddUser(params.Username, params.Password, params.Role, c)
	if e != nil {
		response.ErrorResponse(c, e)
		return
	}
	response.NilResponse(c)

}

// ModifyUser
// @Tags 用户管理
// @Summary 修改用户
// @Description 修改用户
// @Accept  json
// @Produce  json
// @Param id path string true "用户id"
// @Param param body request.UserModifyRequest true "body"
// @Param Auth header string true "Token"
// @Success 200 {object} response.Response "desc"
// @Router /user/{id} [put]
func (mine *userController) ModifyUser(c *gin.Context) {
	log := mine.logger.NewInstance(c)
	id := c.Param("id")
	if id == "" {
		log.Error("id is empty = ")
		response.ParamsErrorResponse(c)
	}
	params := request.UserModifyRequest{}
	e := c.BindJSON(&params)
	if e != nil {
		log.Error("json parse error")
		response.ParamsErrorResponse(c)
		return
	}
	success, _ := govalidator.ValidateStruct(params)
	if !success {
		log.Error("params is empty = ", params)
		response.ParamsErrorResponse(c)
		return
	}

	e = mine.userService.ModifyUser(id, params.Role)
	if e != nil {
		response.ErrorResponse(c, e)
		return
	}

	response.NilResponse(c)

}

// DelUser
// @Tags 用户管理
// @Summary 删除用户
// @Description 删除用户
// @Accept  json
// @Produce  json
// @Param id path string true "用户id"
// @Param Auth header string true "Token"
// @Success 200 {object} response.Response
// @Router /user/{id} [delete]
func (mine *userController) DelUser(c *gin.Context) {
	log := mine.logger.NewInstance(c)
	id := c.Param("id")
	if id == "" {
		log.Error("id is empty = ")
		response.ParamsErrorResponse(c)
	}
	e := mine.userService.DelById(id)
	if e != nil {
		response.ErrorResponse(c, e)
		return
	}
	response.NilResponse(c)
}

// ListUser
// @Tags 用户管理
// @Summary 用户列表
// @Description 用户列表
// @Accept  json
// @Produce  json
// @Param pageIndex query string true "pageIndex"
// @Param pageSize query string true "pageSize"
// @Param username query string true "用户名"
// @Param role query string true "角色"
// @Param Auth header string true "Token"
// @Success 200 {object} response.UserListResponse "desc"
// @Router /user [get]
func (mine *userController) ListUser(c *gin.Context) {
	log := mine.logger.NewInstance(c)
	pageIndexStr := c.Query("pageIndex")
	pageSizeStr := c.Query("pageSize")
	if pageSizeStr == "" || pageIndexStr == "" {
		response.ParamsErrorResponse(c)
		return
	}

	pageIndex, e0 := strconv.ParseInt(pageIndexStr, 10, 64)
	pageSize, e1 := strconv.ParseInt(pageSizeStr, 10, 64)

	if e0 != nil || e1 != nil {
		log.Error("pageSize or pageIndex error")
		response.ParamsErrorResponse(c)
		return
	}

	params := request.UserListRequest{}
	params.PageIndex = pageIndex
	params.PageSize = pageSize
	params.Username = c.Query("username")
	params.Role = enum.RoleEnum(c.Query("role"))

	// check
	success, _ := govalidator.ValidateStruct(params)
	if !success {
		log.Error("params is empty = ", params)
		response.ParamsErrorResponse(c)
		return
	}

	// do service
	data, total, e := mine.userService.ListUser(&params)
	if e != nil {
		response.ParamsErrorResponse(c)
		return
	}

	resp := response.UserListResponse{}

	var users []response.UserDetailResponse

	for _, item := range *data {
		var r response.UserDetailResponse
		r.Id = item.Id.Hex()
		r.Username = item.Username
		r.Role = item.Role
		users = append(users, r)
	}

	resp.Users = &users
	resp.Total = total

	response.DataResponse(c, resp)

}

// DetailUser
// @Tags 用户管理
// @Summary 用户详情
// @Description 用户详情
// @Accept  json
// @Produce  json
// @Param id path string true "用户id"
// @Param Auth header string true "Token"
// @Success 200 {object} response.UserDetailResponse "desc"
// @Router /user/{id} [get]
func (mine *userController) DetailUser(c *gin.Context) {
	// @Success 200 {object} response.Response{data=response.UserDetailResponse}
	log := mine.logger.NewInstance(c)
	id := c.Param("id")
	if id == "" {
		log.Info("id is empty")
		response.ParamsErrorResponse(c)
		return
	}
	res, e := mine.userService.QueryById(id)
	if e != nil {
		response.ErrorResponse(c, e)
		return
	}
	var resp = response.UserDetailResponse{}
	resp.Id = res.Id.Hex()
	resp.Username = res.Username
	resp.Role = res.Role

	response.DataResponse(c, resp)

}
