package controllers

import (
	"log"

	"github.com/Asprilla24/go-ums/app/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1)/UMS?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panicf(err.Error())
	}

	db.DB()
	db.AutoMigrate(&models.AppUsers{})
	DB = db
}
