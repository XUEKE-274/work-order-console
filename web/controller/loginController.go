package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"work-order-console/config"
	"work-order-console/domain/enum/errorCodeEnum"
	"work-order-console/logger"
	"work-order-console/service"
	"work-order-console/web/request"
	"work-order-console/web/response"
)

type LoginControllerApi interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type loginController struct {
	redisService service.RedisServiceApi
	userService service.UserServiceApi
	logger logger.NewLogger
	sessionService service.SessionServiceApi
	config *config.Config
	rsaService service.RsaServiceApi
}


var regLoginController = fx.Provide(func(redisService service.RedisServiceApi,
									     userService service.UserServiceApi,
									     logger logger.NewLogger,
										 sessionService service.SessionServiceApi,
										 config *config.Config,
										 rsaService service.RsaServiceApi) LoginControllerApi{
	return &loginController{
		redisService,
		userService,
		logger,
		sessionService,
		config,
		rsaService,
	}
})

// Login
// @Tags 登陆管理
// @Summary 登陆
// @Description 登陆
// @Accept  json
// @Produce  json
// @Param param body request.LoginRequest true "body"
// @Success 200 {object} response.Response "desc"
// @Router /login [post]
func (mine *loginController)Login(c *gin.Context)  {
	log := mine.logger.NewInstance(c)
	log.Info("login success")
	var params = request.LoginRequest{}
	e := c.BindJSON(&params)
	if e != nil {
		log.Info("json parse error")
		response.ParamsErrorResponse(c)
		return
	}
	success, _ := govalidator.ValidateStruct(params)
	if !success {
		log.Error("request = ", params)
		response.ParamsErrorResponse(c)
		return
	}
	// do login by userService
	userPo, err := mine.userService.GetByUsername(params.Username)
	if err != nil {
		log.Error("dao error ")
		response.ErrorCodeResponse(c, errorCodeEnum.USERNAME_NOT_EXIST)
		return
	}
	// Decrypt
	var pwd string
	if mine.config.Rsa == "ON" {
		pwd = mine.rsaService.Decrypt(params.Password)
	}else {
		pwd = params.Password
	}
	success = mine.userService.MatchCredential(pwd, userPo.EnPwd)
	if success {
		token := mine.sessionService.GenToken(userPo)
		var res = response.LoginSuccessResponse{
			Auth: token,
		}
		response.DataResponse(c, res)
	}else {
		response.ErrorCodeResponse(c, errorCodeEnum.AUTH_FAIL)
	}


}
// Logout
// @Tags 登陆管理
// @Summary 退出登陆
// @Description 退出登陆
// @Accept  json
// @Produce  json
// @Param Auth header string true "Token"
// @Success 200 {object} response.Response "desc"
// @Router /logout [get]
func (mine *loginController)Logout(c *gin.Context)  {
	token := c.GetHeader("Auth")
	if token == ""{
		response.ParamsErrorResponse(c)
		return
	}
	// do logout
	mine.redisService.Del(token)
	response.NilResponse(c)
}
