package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Product struct {
	Id     int    `gorm:"AUTO_INCREMENT"`
	Name   string `gorm:"size:50" form:"name" binding:"required"`
	Price  int    `form:"price" binding:"required"`
	UserId int
	Images []Image `gorm:"foreignkey:ProductId`
}


func (p *Product) AddProduct(db *gorm.DB) error {
	err := db.Create(&p).Error
	return err

}

func (p *Product) DeleteProduct(db *gorm.DB) error {
	err := db.First(&p).Error
	if err != nil {
		return err
	}
	err = db.Delete(&p).Error
	return err
}

func (p *Product) UpdateProduct(db *gorm.DB) error {
	err := db.Save(&p).Error
	return err
}

func GetAllProduct(db *gorm.DB) ([]Product, error) {
	var products []Product
	err := db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *Product) GetProduct(db *gorm.DB) error {
	err := db.Model(&p).Related(&p.Images).Error
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
	}
	return err
}

func (p *Product) CheckProduct(db *gorm.DB) error {
	err := db.First(&p).Error
	return err
}

func (u *User) GetImageProduct(db *gorm.DB) error {
	err := db.Model(&u).Related(&u.Products).Error
	return err
}
