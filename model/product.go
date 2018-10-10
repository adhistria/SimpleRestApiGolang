package model

import (
	"database/sql"
	"fmt"
)

type Product struct {
	Id     int
	Name   string
	Price  int
	UserId int
}

func (p *Product) AddProduct(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO PRODUCTS (NAME, PRICE, USER_ID) Values('%s',%d,%d)", p.Name, p.Price, p.UserId)
	res, err := db.Exec(statement)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	p.Id = int(id)
	return err

}

func (p *Product) DeleteProduct(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE PRODUCTS WHERE ID= %d", p.Id)
	_, err := db.Exec(statement)
	return err
}

func (p *Product) UpdateProduct(db *sql.DB) error{
	statement := fmt.Sprintf("UPDATE USERS SET NAME='%s' Price= %d where id=%d",p.Name,p.Price,p.Id)
	_, err := db.Exec(statement)
	return err
}

func GetAllProduct(db *sql.DB) ([]Product, error){
	statement := fmt.Sprintf("SELECT * FROM PRODUCT")
	var products []Product
	rows, err := db.Query(statement)
	for rows.Next(){
		p:= Product{}
		err := rows.Scan(&p.Id, &p.Name, &p.Price, &p.UserId)
		if err!= nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products,err
}

func (p *Product) GetProduct(db *sql.DB) error{
	statement := fmt.Sprintf("SELECT * FROM USERS WHERE ID = %d",p.Id)
	err := db.QueryRow(statement).Scan(p.Id, p.Name, p.Price, p.UserId)
	return err
}