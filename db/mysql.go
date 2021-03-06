package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"work-order-console/config"
	"work-order-console/domain/entity"
)

var mysql = fx.Provide(func(logger *logrus.Logger, config *config.Config) *gorm.DB {
	var db *gorm.DB
	db, _ = gorm.Open("mysql", config.MysqlUri)
	db.LogMode(true) // debug
	db.SingularTable(true) // 单数
	db.SetLogger(logger)
	logger.Info("mysql table init >>> ")
	db.AutoMigrate(&entity.TemplateEntity{})
	return db
})