package config

import (
	"github.com/imroc/req"
	"go.uber.org/fx"
	"net/http"
	"time"
)

var RpcConfig = fx.Invoke(func() {
	client := &http.Client{}
	client.Transport = &http.Transport{
		MaxIdleConnsPerHost: 500,
	}

	req.SetClient(client)
	req.SetTimeout(5 * time.Second)
})
