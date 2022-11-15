package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func Connect() {
	url := os.Getenv("DB_URL")
	d, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic("Failed to connect db")
	}

	db = d
}

func GetDb() *gorm.DB {
	return db
}
