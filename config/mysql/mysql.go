package mysql

import (
	"fmt"
	"log"
	"web-desa/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGorm(username, password, host, database string, port int) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	db.AutoMigrate(model.Umkm{}, model.Wisata{}, model.Desa{}, model.InfoKegiatan{}, model.User{})

	return db
}