package model

import (
	"github.com/jinzhu/gorm"
)

type Image struct {
	Id        int    `gorm:"AUTO_INCREMENT"`
	Url 	  string 
	ProductId int
}

func (img *Image) AddImage(db *gorm.DB) error {
	err := db.Create(&img).Error
	return err
}

func (img *Image) DeleteImage(db *gorm.DB) error {
	err := db.First(&img).Error
	if err != nil {
		return err
	}
	err = db.Delete(&img).Error
	return err
}

func (img *Image) U(db *gorm.DB) error {
	err := db.First(&img).Error
	if err != nil {
		return err
	}
	err = db.Delete(&img).Error
	return err
}

