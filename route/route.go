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
	// Router := new(middleware.CustomMux)
	// Router.RegisterMiddleware(MiddlewareJWTAuthorization)
	Router = mux.NewRouter()

	// Router.HandleFunc("/user/product/{id:[0-9]+}", controller.GetUserProduct)
	// Router.HandleFunc("/user/{id:[0-9]+}", controller.ReadUser)
	// Router.HandleFunc("/user/{id:[0-9]+}", controller.UpdateUser)
	// Router.HandleFunc("/users", controller.ReadUsers)
	// Router.HandleFunc("/users/{id:[0-9]}", controller.DeleteUser)
	// Router.HandleFunc("/users", controller.AddUser)
	// Router.HandleFunc("/products/{id:[0-9]+}", controller.GetProduct)
	// Router.HandleFunc("/products/{id:[0-9]+}", controller.UpdateProduct)
	// Router.HandleFunc("/products", controller.GetAllProduct)
	// Router.HandleFunc("/products/{id:[0-9]+}", controller.DeleteProduct)
	// Router.HandleFunc("/products", controller.AddProduct)
	Router.HandleFunc("/user/product/{id:[0-9]+}", controller.GetUserProduct).Methods("GET")
	Router.HandleFunc("/user/{id:[0-9]+}", controller.ReadUser).Methods("GET")
	Router.HandleFunc("/user/{id:[0-9]+}", controller.UpdateUser).Methods("PUT")
	Router.HandleFunc("/users", controller.ReadUsers).Methods("GET")
	Router.HandleFunc("/users/{id:[0-9]}", controller.DeleteUser).Methods("DELETE")
	Router.HandleFunc("/users", controller.AddUser).Methods("POST")
	Router.HandleFunc("/products/{id:[0-9]+}", controller.GetProduct).Methods("GET")
	Router.HandleFunc("/products/{id:[0-9]+}", controller.UpdateProduct).Methods("PUT")
	Router.HandleFunc("/products", controller.GetAllProduct).Methods("GET")
	Router.HandleFunc("/products/{id:[0-9]+}", controller.DeleteProduct).Methods("DELETE")
	Router.HandleFunc("/products", controller.AddProduct).Methods("POST")
	Router.HandleFunc("/userinfo", middleware.ValidateMiddleware(controller.GetUserLogin)).Methods("GET")
	Router.HandleFunc("/login", controller.Login).Methods("POST")

	// Router.HandleFunc("/products", controller.AddProduct).Methods("POST")
	// server := new(http.Server)
	// server.Handler = Router
	// server.Addr = ":8000"

	// // fmt.Println("Starting server at", server.Addr)
	// server.ListenAndServe()
}
