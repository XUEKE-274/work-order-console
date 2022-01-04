package request


type LoginRequest struct {
	Username string `json:"username" valid:"required"` //用户名
	Password string `json:"password" valid:"required"` //密码， 使用公钥加密后传输
}