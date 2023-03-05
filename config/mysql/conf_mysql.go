package mysql

import (
	"fmt"
	"gbv2/config/log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func Dsn() string {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.db"),
	)
	return dsn
}

func DBInit() {
	var err error
	dsn := Dsn()
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Panicw("数据库连接失败", err)
	}

	log.Infow("数据库连接成功", "dsn:", dsn)

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	if db == nil {
		log.Errorw("db is nil")
	}
	DB = db
}
