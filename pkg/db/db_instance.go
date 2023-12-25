package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"ha/pkg/config"
)

var DB *gorm.DB

func SetUp(config config.Configuration) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.MySQLUser, config.MySQLPassword, config.MySQLHostName, config.MySQLPort, config.MySQLDB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	db.NamingStrategy = schema.NamingStrategy{SingularTable: true}
	DB = db
	return nil
}
