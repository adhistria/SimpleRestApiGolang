package route

import (

	"rest_api/controller"
	"github.com/gorilla/mux"
)

var Router *mux.Router

func init(){
	Router = mux.NewRouter()
	
	Router.HandleFunc("/user/{id:[0-9]+}", controller.ReadUser).Methods("GET")
	Router.HandleFunc("/user/{id:[0-9]+}", controller.UpdateUser).Methods("PUT")
	Router.HandleFunc("/users", controller.ReadUsers).Methods("GET")
	Router.HandleFunc("/users/{id:[0-9]}", controller.DeleteUser).Methods("DELETE")
	Router.HandleFunc("/user", controller.AddUser).Methods("POST")
}
