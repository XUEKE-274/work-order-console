package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"work-order-console/service"
	"work-order-console/web/response"
)

type SystemControllerApi interface {
	FetchPubKey(c *gin.Context)
	Test(*gin.Context)
}

type systemController struct {
	rsaService service.RsaServiceApi
	logService service.LockServiceApi
}



var regSystemController = fx.Provide(func(rsaService service.RsaServiceApi, lockService service.LockServiceApi) SystemControllerApi {
	return &systemController{
		rsaService,
		lockService,
	}
})

// FetchPubKey
// @Tags 系统关系
// @Summary 获取公钥
// @Description 获取公钥
// @Accept  json
// @Produce  json
// @Success 200 {object} response.PubKeyResponse "desc"
// @Router /system/pubKey [get]
func (mine *systemController) FetchPubKey(c *gin.Context) {
	pubStr := mine.rsaService.GetPubKey()
	var data response.PubKeyResponse
	data.PubKey = pubStr
	response.DataResponse(c, data)
}

func (mine *systemController)Test(c *gin.Context)  {
	r := mine.logService.TryLock("test")
	data := make(map[string]interface{})
	data["lock"] = r
	response.DataResponse(c, data)


}