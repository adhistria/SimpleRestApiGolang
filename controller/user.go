package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"rest_api/database"
	"github.com/gorilla/mux"
	"rest_api/model"
	"rest_api/respond"
)

func ReadUser(w http.ResponseWriter, r *http.Request) {
	// read params
	vars := mux.Vars(r)
	// convert string to integer
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid User Id")
		return
	}
	u := model.User{Id: id}
	err = u.GetUser(database.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respond.RespondWithError(w, http.StatusNotFound, "User Not Found")
		default:
			respond.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, u)
	return
}

func ReadUsers(w http.ResponseWriter, r *http.Request) {
	users, err := model.GetUsers(database.DB)
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, users)
}


func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid User Id")
		return
	}
	u := model.User{Id: id}
	err = u.DeleteUser(database.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respond.RespondWithError(w, http.StatusBadRequest, "User Not Found")
		default:
			respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		}
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, "User deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid User Id")
	}

	var u model.User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&u)
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid Request Payload")
		return
	}
	defer r.Body.Close()
	u.Id = id
	err = u.UpdateUser(database.DB)
	if err != nil {
		fmt.Println("error not nil")
		switch err {
		case sql.ErrNoRows:
			respond.RespondWithError(w, http.StatusNotFound, "User Not Found")
		default:
			respond.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, u)

	// Kalo pake form
	// r.ParseForm()
	// for key, value := range r.Form {
	// 	fmt.Printf("%s = %s\n", key, value)
	// }
	// fmt.Printf("Name => %s\n", r.FormValue("Name"))
	// fmt.Printf("Age => %s\n", r.FormValue("Age"))
	// // age := r.FormValue("Age")
	// var u User
	// u.Id = id
	// u.Name = r.FormValue("Name")
	// u.Age, err = strconv.Atoi(r.FormValue("Age"))
	// if err != nil {
	// 	RespondWithError(w, http.StatusBadRequest, "Invalid Age Person")
	// 	return
	// }
	// err = u.UpdateUser(database.DB)
	// if err != nil {
	// 	switch err {
	// 	case sql.ErrNoRows:
	// 		RespondWithError(w, http.StatusNotFound, "User Not Found")
	// 	default:
	// 		RespondWithError(w, http.StatusBadRequest, err.Error())
	// 		return
	// 	}
	// }
	// RespondWithJSON(w, http.StatusOK, u)
	//

}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var u model.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
	}
	defer r.Body.Close()
	err = u.AddUser(database.DB)
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid payload")
	}
	respond.RespondWithJSON(w, http.StatusOK, u)
}

