package model

import (
	// "database/sql"
	"golang.org/x/crypto/bcrypt"

	"fmt"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id       int    `gorm:"AUTO_INCREMENT"`
	Username string `gorm:"size:50"`
	Email    string `gorm:"size:50"`
	Age      int
	Password string
	Name     string    `gorm:"size:50"`
	Products []Product `gorm:"foreignkey:User_Id`
}

type UserLogin struct {
	Username string
	Password string
}

type Claims struct {
	jwt.StandardClaims
	Username string `json:"Username"`
	UserId   int    `json:"User_Id"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

func (u *User) GetUser(db *gorm.DB) error {
	// statement := fmt.Sprintf("SELECT * FROM USERS WHERE ID=%d", u.Id)
	// err:= db.QueryRow(statement).Scan(&u.Id, &u.Name, &u.Age)
	// if err!= nil {
	// 	return nil, err
	// }
	err := db.First(&u).Error
	return err

	// f := map[string]interface{}{
	// 	"Id":   u.Id,
	// 	"Name": u.Name,
	// 	"Age":  u.Age,
	// }

}

func GetUsers(db *gorm.DB) ([]User, error) {
	// statement := fmt.Sprintf("SELECT * FROM USERS")
	// rows, err := db.Query(statement)
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()
	// users := []User{}
	// for rows.Next() {
	// 	var u User
	// 	err = rows.Scan(&u.Id, &u.Name, &u.Age)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	users = append(users, u)
	// }
	// return users, nil
	var users []User
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) UpdateUser(db *gorm.DB) error {
	// statement := fmt.Sprintf("UPDATE USERS  SET Name='%s', Age=%d WHERE ID=%d", u.Name, u.Age, u.Id)
	// fmt.Println("db update user")
	// // statement := fmt.Sprintf("UPDATE users SET name='%s', age=%d WHERE id=%d", u.Name, u.Age, u.ID)

	// _, err := db.Exec(statement)
	// return err
	db.First(&u)
	err := db.Save(&u).Error
	return err

}

func (u *User) DeleteUser(db *gorm.DB) error {
	// statement := fmt.Sprintf("DELETE USERS WHERE ID=%d", u.Id)
	// _, err := db.Exec(statement)
	// return err
	err := db.First(&u).Error
	if err != nil {
		return err
	}
	err = db.Delete(&u).Error
	return err
}

func (u *User) AddUser(db *gorm.DB) error {
	// statement := fmt.Sprintf("INSERT INTO USERS (name,age) values('%s',%d)", u.Name, u.Age)
	// res, err := db.Exec(statement)
	// if err != nil {
	// 	return err
	// } else {
	// 	id, err := res.LastInsertId()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	u.Id = int(id)
	// 	return err
	// }
	// newPassword := hashPassword(u.Password)
	newPassword := hashPassword("password")
	fmt.Println(u)
	u.Password = string(newPassword)
	fmt.Println(u)
	db.NewRecord(u)
	err := db.Create(&u).Error
	return err
}

func (u *User) GetUserProduct(db *gorm.DB) error {
	// statement := fmt.Sprintf("SELECT * FROM USERS.ID INNER JOIN PRODUCTS.ID ON USERS.ID=PRODUCTS.ID WHERE USER.ID=%d",u.Id)
	// statement1 := fmt.Sprintf("SELECT * FROM USERS WHERE ID=%d", u.Id)
	// err := db.QueryRow(statement1).Scan(&u.Id, &u.Name, &u.Age)
	// if err != nil {
	// 	return err
	// }
	// statement2 := fmt.Sprintf("SELECT products.id, products.name, products.price FROM products INNER JOIN users ON USERS.ID=PRODUCTS.USER_ID WHERE users.id=%d", u.Id)
	// rows, err := db.Query(statement2)
	// for rows.Next() {
	// 	var product Product
	// 	err := rows.Scan(&product.Id, &product.Name, &product.Price)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	u.Products = append(u.Products, product)
	// }

	// m := f.(map[string]interface{})
	// var products []Product

	// res := db.Find(&u)
	// if len(res.GetErrors())>0 {
	// 	return res.GetErrors()
	// }
	// statement := fmt.Sprintf("USER_ID = %d",u.Id)
	// resProduct := db.Where(statement).Find(&u.Products)
	err := db.Model(&u).Related(&u.Products).Error

	// res := db.Model(&u).Find(products)
	// fmt.Println(res.GetErrors())
	return err
}

func (u *User) Login(db *gorm.DB) error {
	// byteHash := hashPassword(u.Password)
	password := u.Password
	// fmt.Println("hash password", byteHash)
	err := db.Where("username = ?", u.Username).First(&u).Error
	// fmt.Println(u)
	if err != nil {
		fmt.Println("ga dapat datanya")
		return err
	}
	// fmt.Println(byteHash, u.Password)
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	// fmt.Println(byteHash, []byte(u.Password))
	if err != nil {
		fmt.Println("ga sama")
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
