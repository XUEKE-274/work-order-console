package config

import (
	"go.uber.org/fx"
)






var Model = fx.Options(
	YmlConfig,
	RpcConfig,
)

