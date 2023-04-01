package router

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/purvakoushik/mongo_DB_API/controller"
)

// function Routers to route the network traffic to differnt paths according to parameters and methods.
func Routers() *mux.Router {

	fmt.Println("The API routers for mongoDB")

	newRouter := mux.NewRouter()

	//Routers for controllers that will take help form the mongoDB helper for CRUD operations.
	newRouter.HandleFunc("/api", controller.ServeHome).Methods("GET")
	newRouter.HandleFunc("/api/movies", controller.GetAllMovie).Methods("GET")
	newRouter.HandleFunc("/api/movie", controller.CreateOneMovie).Methods("POST")
	newRouter.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	newRouter.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	newRouter.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovies).Methods("DELETE")

	return newRouter
}
