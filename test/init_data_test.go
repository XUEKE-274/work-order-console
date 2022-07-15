package test

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
	"work-order-console/utils"
)

func TestInit(t *testing.T)  {
	password := "ZXC@@123"

	salt := utils.NewUuid()
	target := password + salt
	h := sha256.New()
	h.Write([]byte(target))
	sum := h.Sum(nil)
	userPwdHash := hex.EncodeToString(sum)
	r := userPwdHash + salt
	println(r)
}
