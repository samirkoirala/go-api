package main

import (
	"log"
	"net/http"

	"github.com/Sulav-Adhikari/gouser/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterUserStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("0.0.0.0:9010", r))

}
