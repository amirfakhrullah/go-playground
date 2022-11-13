package config

import (
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	url := os.Getenv("RAILWAY_MYSQL")
	d, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic("failed to connect db")
	}

	db = d
}

func GetDb() *gorm.DB {
	return db
}
