package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"scry/global"
	"time"
)

// InitMysql
//  @Description: 初始化数据库连接
//  @return *gorm.DB
//  @author	wujingfeng 2022-07-05 14:42:46
func InitMysql() *gorm.DB {
	m := global.Config.Mysql
	dsn := m.DSN()
	mysqlConfig := mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 191,
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sdb, err := db.DB()
	sdb.SetMaxIdleConns(10)
	sdb.SetMaxOpenConns(100)
	sdb.SetConnMaxLifetime(30 * time.Second)
	return db
}
