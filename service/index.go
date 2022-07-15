package service

import "go.uber.org/fx"

var Model = fx.Options(
	redisServiceReg,
	regUserService,
	regTicketService,
	regSessionService,
	regRsaService,
	regTemplateService,
	regGrpcService,
	regLockService,
)