package database

import (
	// "fmt"
	"log"

	// "database/sql"
	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"rest_api/model"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:@/golang_new")
	DB.DropTable(&model.Product{})
	DB.DropTable(&model.User{})
	user := model.User{Username:"adhisatria",Password:"password",Age:15,Email:"adhistria1@gmail.com",Name:"Adhi Satria"}


	DB.AutoMigrate(&model.Product{})
	DB.AutoMigrate(&model.User{})
	DB.NewRecord(user) 
	DB.Create(&user)
	// DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_new")
	if err != nil {
		log.Fatal(err)
	}
}
