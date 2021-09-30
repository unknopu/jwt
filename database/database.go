package database

import (
	"jwt/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func SetUpDB(){
	database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil{
		panic("failed to create database")
	}

	database.AutoMigrate(&model.User{})
	db = database
}
