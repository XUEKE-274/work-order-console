package db

import "go.uber.org/fx"

var Model = fx.Options(
	mongodb,
	redisDb,
	mysql,
)