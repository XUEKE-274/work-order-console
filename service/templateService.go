package service

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/fx"
	"work-order-console/domain/entity"
	"work-order-console/logger"

)

type TemplateServiceApi interface {
	QueryAll() *[]entity.TemplateEntity
}

type templateService struct {
	logger.NewLogger
	*gorm.DB
}

func (p *templateService)QueryAll() *[]entity.TemplateEntity {
	var res []entity.TemplateEntity
	p.Find(&res)
	return &res
}

var regTemplateService = fx.Provide(func(logger logger.NewLogger, db *gorm.DB) TemplateServiceApi {
	return &templateService{
		logger,
		db,
	}
})
