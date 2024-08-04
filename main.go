package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/oneitworld-demo-crud-api-go/commons"
	"github.com/oneitworld-demo-crud-api-go/routes"
)

const SERVER_PORT string = ":9090"

func main() {
	commons.Migrate()

	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)

	// Asegúrate de que el directorio uploads exista

	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		fmt.Println("Error al crear el directorio uploads:", err)
		return
	}

	server := http.Server{
		Addr:    SERVER_PORT,
		Handler: router,
	}

	log.Println("AI World Talent | HTTP Server READY on Port " + SERVER_PORT)
	log.Println(server.ListenAndServe())

}
