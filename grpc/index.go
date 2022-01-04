package grpc

import (
	"go.uber.org/fx"
	"work-order-console/grpc/template"
)

var Model = fx.Options(
	template.RegServerTemplate,
)