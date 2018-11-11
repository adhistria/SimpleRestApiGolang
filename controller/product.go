package controller

import (
	"fmt"
	// "errors"
	"rest_api/model"
	// "database/sql"
	"encoding/json"
	"net/http"
	"rest_api/database"
	"rest_api/respond"
	"strconv"
	// "io/ioutil"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	decoder := json.NewDecoder(r.Body)
	var product model.Product
	err := decoder.Decode(&product)
	product.User_Id = int(userInfo["User_Id"].(float64))
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid Payload Request")
		return
	}
	err = product.AddProduct(database.DB)
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, "Fail Add New Product")
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid Product Id")
	}

	p := model.Product{Id: id}
	// var errors []error
	// res, err := p.DeleteProduct(database.DB)
	err = p.DeleteProduct(database.DB)
	// fmt.Println(res)
	if err != nil {
		// for _, err  := range errors {
		// 	arr_string_err = append(arr_string_err, err.Error())
		// }
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, p)
	// if err{
	// 	switch err{
	// 	case sql.ErrNoRows:
	// 		respond.RespondWithError(w, http.StatusBadRequest, "Product Not Found")
	// 	default:
	// 		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
	// 	}
	// 	return
	// }
	// var make(map[]string interface{})
	// respond.RespondWithJSON(w, http.StatusOK, "Delete Success")

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	fmt.Println(id)
	if err != nil {
		// arr_string_err = append(arr_string_err,)
		fmt.Println(err.Error())
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid Product Id")
		return
	}
	p := model.Product{Id: id}	
	json.NewDecoder(r.Body).Decode(&p)
	p.User_Id = int(userInfo["User_Id"].(float64))
	fmt.Println(p)
	err = p.UpdateProduct(database.DB)

	json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, p)
}

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	products, err := model.GetAllProduct(database.DB)
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// arr_string_err = append(arr_string_err,)
		respond.RespondWithError(w, http.StatusBadRequest, "Product Id")
		return
	}
	p := model.Product{Id: id}
	err = p.GetProduct(database.DB)
	fmt.Println("sebenernya udah masuk get product")
	if err != nil {
		// arr_string_err = append(arr_string_err,)
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, p)
	// if err!= nil {
	// 	switch err{
	// 	case sql.ErrNoRows:
	// 		respond.RespondWithError(w, http.StatusBadRequest, "Product Not Found")
	// 	default:
	// 		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
	// 	}
	// 	return
	// }
	// respond.RespondWithJSON(w, http.StatusOK, p)

}
