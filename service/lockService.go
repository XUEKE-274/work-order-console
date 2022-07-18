package service

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/fx"
	"work-order-console/domain/entity"
)

type LockServiceApi interface {
	TryLock(string) bool
}

type lockService struct {
	*gorm.DB
}


var regLockService = fx.Provide(func(db *gorm.DB) LockServiceApi{
	return &lockService{
		db,
	}
})

func (p *lockService)TryLock(resource string) bool {
	if resource == "" {
		return false
	}
	// all select
	vs := make([]entity.VersionEntity, 0)
	p.Find(&vs)
	// select
	var v entity.VersionEntity
	p.Where("resource = ?", resource).Find(&v)
	version := v.Version
	// do update
	res := p.Model(&entity.VersionEntity{}).
		Where("resource = ? and version = ?", resource, version).
		Update("version" , version + 1)
	return res.RowsAffected > 0
}
