package rpc

import "go.uber.org/fx"

var Model = fx.Options(
	accountRpcReg,
)
