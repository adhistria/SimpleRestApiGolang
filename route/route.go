package route

import (
	// "net/http"
	"rest_api/controller"

	"rest_api/middleware"
	// "time"

	"github.com/gorilla/mux"
)

var Router *mux.Router

type M map[string]interface{}

func init() {
	Router = mux.NewRouter()
	Router.HandleFunc("/user/product/{id:[0-9]+}", controller.GetUserProduct).Methods("GET")
	Router.HandleFunc("/user/{id:[0-9]+}", controller.ReadUser).Methods("GET")
	Router.HandleFunc("/user/{id:[0-9]+}", controller.UpdateUser).Methods("PUT")
	Router.HandleFunc("/users", controller.ReadUsers).Methods("GET")
	Router.HandleFunc("/users/{id:[0-9]}", controller.DeleteUser).Methods("DELETE")
	Router.HandleFunc("/users", controller.AddUser).Methods("POST")
	Router.HandleFunc("/products/{id:[0-9]+}", controller.GetProduct).Methods("GET")
	Router.HandleFunc("/products/{id:[0-9]+}", middleware.ValidateMiddleware(controller.UpdateProduct)).Methods("PUT")
	Router.HandleFunc("/products", controller.GetAllProduct).Methods("GET")
	Router.HandleFunc("/products/{id:[0-9]+}", middleware.ValidateMiddleware(controller.DeleteProduct)).Methods("DELETE")
	Router.HandleFunc("/products", middleware.ValidateMiddleware(controller.AddProduct)).Methods("POST")
	Router.HandleFunc("/userinfo", middleware.ValidateMiddleware(controller.GetUserLogin)).Methods("GET")
	Router.HandleFunc("/login", controller.Login).Methods("POST")
}
