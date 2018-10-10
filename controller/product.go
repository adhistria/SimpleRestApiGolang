package controller

import (
	"rest_api/model"
	"database/sql"
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
	if err!= nil {
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid Payload Request")
		return
	}
	
	err = product.AddProduct(database.DB)
	if err!= nil {
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respond.RespondWithJSON(w,http.StatusOK, product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!= nil {
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid Product Id")
	}

	p := model.Product{Id:id}
	err = p.DeleteProduct(database.DB)
	if err!= nil{
		switch err{
		case sql.ErrNoRows:
			respond.RespondWithError(w, http.StatusBadRequest, "Product Not Found")
		default:
			respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		}
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, "Delete Success")
}

func UpdateProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!= nil {
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid Payload")
		return
	}
	p := model.Product{Id:id}
	err = p.UpdateProduct(database.DB)
	if err!= nil {
		switch err{
		case sql.ErrNoRows:
			respond.RespondWithError(w, http.StatusBadRequest, "Product Not Found")
		default:
			respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		}
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, p)
}

func GetAllProduct(w http.ResponseWriter, r *http.Request){
	products, err := model.GetAllProduct(database.DB)	
	if err!= nil {
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, products)
}

func GetProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!= nil {
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid Product Id")
		return
	}
	p := model.Product{Id:id}
	err = p.GetProduct(database.DB)
	if err!= nil {
		switch err{
		case sql.ErrNoRows:
			respond.RespondWithError(w, http.StatusBadRequest, "Product Not Found")
		default:
			respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		}
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, p)
}