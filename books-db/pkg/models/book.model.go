package models

import (
	"github.com/jinzhu/gorm"
	"books-db/pkg/config"
)

var db * gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication:"`
}

func init() {
	config.Connect()
}