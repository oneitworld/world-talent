package routes

import (
	"github.com/gorilla/mux"
	"github.com/oneitworld-demo-crud-api-go/controllers"
)

func SetPersonaRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/persona/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/save", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/delete", controllers.GetAll).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.GetByID).Methods("GET")
	subRoute.HandleFunc("/health", controllers.Health).Methods("GET")

}
