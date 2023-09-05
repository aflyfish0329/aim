package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/tinode?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug()

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	sqlDb.SetMaxIdleConns(4)
	sqlDb.SetMaxOpenConns(100)

	return db
}
