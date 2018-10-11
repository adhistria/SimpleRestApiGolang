package controller

import (
	// "fmt"
	// "errors"
	"rest_api/model"
	// "database/sql"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"rest_api/respond"
	"rest_api/database"
	"strconv"
)



func AddProduct(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var product model.Product
	err := decoder.Decode(&product)
	var arr_string_err []string
	if err!= nil {
		// var arr_string []string
		arr_string_err = append(arr_string_err, "Invalid Payload Request")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	fail := product.AddProduct(database.DB)
	if fail {
		arr_string_err = append(arr_string_err, "Fail Add New Product")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w,http.StatusOK, product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	var arr_string_err []string
	if err!= nil {
		arr_string_err = append(arr_string_err, "Invalid Product Id")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
	}

	p := model.Product{Id:id}
	// var errors []error
	res, errors := p.DeleteProduct(database.DB)
	// fmt.Println(res)
	if len(errors)>0 {
		for _, err  := range errors {
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, res)
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

func UpdateProduct(w http.ResponseWriter, r *http.Request){
	var arr_string_err []string
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!= nil {
		arr_string_err = append(arr_string_err,"Invalid Payload")
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	p := model.Product{Id:id}
	errors := p.UpdateProduct(database.DB)
	if len(errors)>0 {
		for _,err := range errors{
			arr_string_err = append(arr_string_err,err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, p)
}

func GetAllProduct(w http.ResponseWriter, r *http.Request){
	products, errors := model.GetAllProduct(database.DB)	
	var arr_string_err []string
	if len(errors)>0 {
		for _,err := range errors{
			arr_string_err = append(arr_string_err, err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, products)
}

func GetProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	var arrStringErr []string
	if err!= nil {
		arrStringErr = append(arrStringErr,err.Error())
		respond.RespondWithError(w, http.StatusBadRequest, arrStringErr)
		return
	}
	p := model.Product{Id:id}
	errors := p.GetProduct(database.DB)
	if len(errors)>0 {
		for _,err := range errors{
			arrStringErr = append(arrStringErr,err.Error())
		}
		respond.RespondWithError(w, http.StatusBadRequest, arrStringErr)
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