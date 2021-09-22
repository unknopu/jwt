package database

import (
	"jwt/model"

	// "gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func SetUpDB(){
	// dsn := "host=localhost user=postgre password=password dbname=gorm port=5432"
	// database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	mydb, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil{
		panic("failed to create database")
	}

	mydb.AutoMigrate(&model.User{})
	mydb.AutoMigrate(&model.Product{})

	db = mydb
}
