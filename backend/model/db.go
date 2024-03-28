package model

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB
var err error

func init() {
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&User{})
}