package controller

import (
	"go.uber.org/fx"
)

var Model = fx.Options(
	regSystemController,
	regLoginController,
	regUserController,
	regTicketController,
	regTemplateController,
)