package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := "root:@tcp(127.0.0.1:3306)/gotodo?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	AutoMigrate()
	return nil
}

func CloseDB() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			panic(err) // Handle the error properly in your application
		}
		sqlDB.Close()
	}
}

func AutoMigrate() {
	DB.AutoMigrate(&User{}, &Task{}) // Pass your models here
}
