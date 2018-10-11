package model

// import (
// 	"database/sql"
// 	"fmt"
// )

// type User struct {
// 	Id       int
// 	Name     string
// 	Age      int
// 	Products []Product
// }

// func (u *User) GetUser(db *sql.DB) (interface{},error) {
// 	statement := fmt.Sprintf("SELECT * FROM USERS WHERE ID=%d", u.Id)
// 	err:= db.QueryRow(statement).Scan(&u.Id, &u.Name, &u.Age)
// 	if err!= nil {
// 		return nil, err
// 	}
// 	f := map[string]interface{}{
// 		"Id": u.Id,
// 		"Name":  u.Name,
// 		"Age": u.Age,
// 	}
	
// 	return f, nil

// }

// func GetUsers(db *sql.DB) ([]User, error) {
// 	statement := fmt.Sprintf("SELECT * FROM USERS")
// 	rows, err := db.Query(statement)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	users := []User{}
// 	for rows.Next() {
// 		var u User
// 		err = rows.Scan(&u.Id, &u.Name, &u.Age)
// 		if err != nil {
// 			return nil, err
// 		}
// 		users = append(users, u)
// 	}
// 	return users, nil
// }

// func (u *User) UpdateUser(db *sql.DB) error {
// 	statement := fmt.Sprintf("UPDATE USERS  SET Name='%s', Age=%d WHERE ID=%d", u.Name, u.Age, u.Id)
// 	fmt.Println("db update user")
// 	// statement := fmt.Sprintf("UPDATE users SET name='%s', age=%d WHERE id=%d", u.Name, u.Age, u.ID)

// 	_, err := db.Exec(statement)
// 	return err
// }

// func (u *User) DeleteUser(db *sql.DB) error {
// 	statement := fmt.Sprintf("DELETE USERS WHERE ID=%d", u.Id)
// 	_, err := db.Exec(statement)
// 	return err
// }

// func (u *User) AddUser(db *sql.DB) error {
// 	statement := fmt.Sprintf("INSERT INTO USERS (name,age) values('%s',%d)", u.Name, u.Age)
// 	res, err := db.Exec(statement)
// 	if err != nil {
// 		return err
// 	} else {
// 		id, err := res.LastInsertId()
// 		if err != nil {
// 			return err
// 		}
// 		u.Id = int(id)
// 		return err
// 	}
// }

// func (u *User) GetUserProduct(db *sql.DB) error {
// 	// statement := fmt.Sprintf("SELECT * FROM USERS.ID INNER JOIN PRODUCTS.ID ON USERS.ID=PRODUCTS.ID WHERE USER.ID=%d",u.Id)
// 	statement1 := fmt.Sprintf("SELECT * FROM USERS WHERE ID=%d", u.Id)
// 	err := db.QueryRow(statement1).Scan(&u.Id, &u.Name, &u.Age)
// 	if err != nil {
// 		return err
// 	}
// 	statement2 := fmt.Sprintf("SELECT products.id, products.name, products.price FROM products INNER JOIN users ON USERS.ID=PRODUCTS.USER_ID WHERE users.id=%d", u.Id)
// 	rows, err := db.Query(statement2)
// 	for rows.Next() {
// 		var product Product
// 		err := rows.Scan(&product.Id, &product.Name, &product.Price)
// 		if err != nil {
// 			return err
// 		}
// 		u.Products = append(u.Products, product)
// 	}
	
// 	// m := f.(map[string]interface{})


// 	return err
// }
