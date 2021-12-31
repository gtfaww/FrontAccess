package service

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var MysqlClient *gorm.DB

func init() {
	MysqlClient = initDB()
}

func initDB() *gorm.DB {

	dsn := "vcom:vcomvcom@tcp(192.168.166.103:3306)/device_manager_111_all_tcp?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetConnMaxIdleTime(time.Hour)
	sqlDB.SetMaxOpenConns(500)

	return db

}
