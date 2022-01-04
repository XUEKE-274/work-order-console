package router

import "go.uber.org/fx"

const (
	serverName = "/work-order-console"
	api = "/api"
	version = "/v1"
	apiPrefix = serverName + api + version

	ORDER = apiPrefix + "/order"
	KF = apiPrefix + "/kf"
	QT = apiPrefix + "qt"

)



var Model =fx.Options(
	regSwaggerRouter,
	regSystemRouter,
	regProbeRouter,
	regLoginRouter,
	regUserRouter,
	regTicketRouter,
	regTemplateRouter,
	startGin,
)
