package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oneitworld-demo-crud-api-go/commons"
	"github.com/oneitworld-demo-crud-api-go/routes"
)

const SERVER_PORT string = ":9090"

func main() {
	commons.Migrate()

	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)

	server := http.Server{
		Addr:    SERVER_PORT,
		Handler: router,
	}

	log.Println("AI World Talent | HTTP Server READY on Port " + SERVER_PORT)
	log.Println(server.ListenAndServe())

}
