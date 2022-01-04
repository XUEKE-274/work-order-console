package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"work-order-console/domain/enum"
	"work-order-console/domain/enum/errorCodeEnum"
	"work-order-console/exception"
)

type Response struct {
	Message string          `json:"message"` // 提示信息
	Code    string          `json:"code"`    // 提示码
	Data    interface{}     `json:"data"`    // 数据
	RequestId  string          `json:"requestId"` // 请求id
}

func ErrorResponse(c *gin.Context, e error)  {
	code := exception.Code(e)
	ErrorCodeResponse(c, code)
}

func ErrorCodeResponse(c *gin.Context, code enum.ErrorCodeEnum)  {
	requestId := c.Value("requestId").(string)
	r := &Response{
		Message: errorCodeEnum.GetValueByCode(code),
		Code: string(code),
		Data: "",
		RequestId: requestId,
	}
	c.JSON(http.StatusOK, r)
}

func NilResponse(c *gin.Context)  {
	DataResponse(c, "")
}
func DataResponse(c *gin.Context, data interface{}) {
	requestId := c.Value("requestId").(string)
	r := &Response{
		Message: "操作成功",
		Code: "SUCCESS",
		Data: data,
		RequestId: requestId,
	}
	c.JSON(http.StatusOK, r)
}

func ParamsErrorResponse(c *gin.Context) {
	ErrorCodeResponse(c, errorCodeEnum.PARAMS_ERROR)
}