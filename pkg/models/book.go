package models

import (
	"github.com/jinzhu/gorm"
	"github.com/raudhra/book-management-system/pkg/config"
)

var db *gorm.DB

type book struct {
	gorm.Model
	Name        string `gorm:"" json"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.getDB()
	db.AutoMigrate(&Book{})
}
