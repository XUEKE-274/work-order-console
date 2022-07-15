package service

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/fx"
)

type LockServiceApi interface {
	TryLock() bool
}

type lockService struct {
	*gorm.DB
}


var regLockService = fx.Provide(func(db *gorm.DB) LockServiceApi{
	return &lockService{
		db,
	}
})

func (p *lockService)TryLock() bool {
	return false
}
