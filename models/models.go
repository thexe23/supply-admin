package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"supply-admin/conf"
)

var db *gorm.DB

func Setup() {
	var err error
	var dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DatabaseSetting.User,
		conf.DatabaseSetting.Password,
		conf.DatabaseSetting.Host,
		conf.DatabaseSetting.Name)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("model set up err: %s", err)
	}
}
