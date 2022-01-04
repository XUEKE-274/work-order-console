package service

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"go.uber.org/fx"
)


var pubKey []byte
var priKey []byte



type RsaServiceApi interface {
	GetPubKey() string
	GetPriKey() string
	Encrypt(source string) string
	Decrypt(source string) string
}

type rsaService struct {

}


var regRsaService = fx.Provide(func() RsaServiceApi{
	privateKey, e := rsa.GenerateKey(rand.Reader, 1024)
	if e != nil {
		panic("rsa.GenerateKey error"+ e.Error())
	}
	X509PublicKey, _ := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)

	pubKey = X509PublicKey
	priKey = x509PrivateKey
	return &rsaService{}
})




func (mine *rsaService) GetPubKey() string {
	return hex.EncodeToString(pubKey)
}

func (mine *rsaService) GetPriKey() string {
	return hex.EncodeToString(priKey)
}

func (mine *rsaService) Encrypt(source string) string {
	pubInterface, _ := x509.ParsePKIXPublicKey(pubKey)
	pub := pubInterface.(*rsa.PublicKey)
	r, _ := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(source))
	return hex.EncodeToString(r)
}

func (mine *rsaService) Decrypt(source string) string {
	pri, _ := x509.ParsePKCS1PrivateKey(priKey)
	r,_ := hex.DecodeString(source)
	t ,_ := rsa.DecryptPKCS1v15(rand.Reader, pri, r)
	return string(t)
}