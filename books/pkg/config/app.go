package config

import (
	"fmt"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func getDbUrl() string {
	url := os.Getenv("RAILWAY_MYSQL")
	fmt.Print("url:", url)
	return url
}

func Connect() {
	url := getDbUrl()
	d, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic("failed to connect db")
	}

	db = d
}

func GetDb() *gorm.DB {
	return db
}
