package controllers

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"eyecandy/app/models"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dbInfo, _ := revel.Config.String("db.info")
	db, err := gorm.Open("mysql", dbInfo)
	if err != nil {
		log.Panicf("Failed gorm.Open: %v\n", err)
	}
	db.DB()
	db.AutoMigrate(&models.Gif{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Tag{})
	db.AutoMigrate(&models.Tagging{})
	
	DB=db
}
