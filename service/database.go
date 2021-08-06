package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go-web-edu/model"
)

var DB *gorm.DB

// 初始化数据库
func InitDB() *gorm.DB {
	driverName := viper.GetString("database.drivername")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	database := viper.GetString("database.database")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	charset := viper.GetString("database.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username, password, host, port, database, charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database,err :" + err.Error())
	}

	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Article{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB.Debug()
}
