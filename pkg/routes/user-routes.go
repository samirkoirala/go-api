package routes

import (
	"net/http" // Add this import

	"github.com/Sulav-Adhikari/gouser/pkg/controllers"
	"github.com/Sulav-Adhikari/gouser/pkg/middleware"
	"github.com/gorilla/mux"
)

var RegisterUserStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/register/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/", controllers.GetAllUser).Methods("GET")
	router.HandleFunc("/signin/", controllers.SignIn).Methods("POST")

	router.Handle("/update-username", middleware.JWTAuthMiddleware(http.HandlerFunc(controllers.UpdateUsername))).Methods("PUT")
}
