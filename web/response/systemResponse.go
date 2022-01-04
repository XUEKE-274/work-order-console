package response


type PubKeyResponse struct{
	PubKey string `json:"pubKey"` //公钥， X509 HEX 格式
}
