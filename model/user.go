package model

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id       int    `gorm:"AUTO_INCREMENT"`
	Username string `gorm:"size:50"`
	Email    string `gorm:"size:50"`
	Password string
	Name     string    `gorm:"size:50"`
	Products []Product `gorm:"foreignkey:UserId`
}

type UserLogin struct {
	Username string
	Name     string
	Token    string `json:"token"`
}

type Claims struct {
	jwt.StandardClaims
	Username string `json:"Username"`
	UserId   int    `json:"UserId"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

func (u *User) GetUser(db *gorm.DB) error {
	err := db.First(&u).Error
	return err
}

func GetUsers(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) UpdateUser(db *gorm.DB) error {
	db.First(&u)
	err := db.Save(&u).Error
	return err

}

func (u *User) DeleteUser(db *gorm.DB) error {
	err := db.First(&u).Error
	if err != nil {
		return err
	}
	err = db.Delete(&u).Error
	return err
}

func (u *User) AddUser(db *gorm.DB) error {
	newPassword := hashPassword(u.Password)
	u.Password = string(newPassword)
	db.NewRecord(u)
	err := db.Create(&u).Error
	return err
}

func (u *User) GetUserProduct(db *gorm.DB) error {
	err := db.Model(&u).Related(&u.Products).Error
	return err
}

func (u *User) Login(db *gorm.DB) error {
	password := u.Password
	err := db.Where("username = ?", u.Username).First(&u).Error
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return err
}

func hashPassword(password string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return hash
}
