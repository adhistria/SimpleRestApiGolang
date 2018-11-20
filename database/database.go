package database

import (
	"log"
	"omahihrom/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:@/golang_new")
	DB.DropTableIfExists(&model.Product{})
	DB.DropTableIfExists(&model.User{})
	DB.DropTableIfExists(&model.Image{})
	DB.AutoMigrate(&model.Product{})
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Image{})
	if err != nil {
		log.Fatal(err)
	}
}
