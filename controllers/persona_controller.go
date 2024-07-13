package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/oneitworld-demo-crud-api-go/commons"
	"github.com/oneitworld-demo-crud-api-go/models"
)

type ServiceStatus struct {
	DBStatus string
}

func Health(writer http.ResponseWriter, request *http.Request) {
	db := commons.GetConnection()
	defer db.Close()

	// Check success connection to Database
	error := db.DB().Ping()
	if error != nil {
		commons.SendError(writer, http.StatusInternalServerError)
		panic(error.Error())

	}

	serviceStatus := ServiceStatus{
		DBStatus: "Healthy",
	}

	json, _ := json.Marshal(serviceStatus)

	commons.SendResponse(writer, http.StatusOK, json)

}

func GetAll(writer http.ResponseWriter, request *http.Request) {
	personas := []models.Persona{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&personas)
	json, _ := json.Marshal(personas)
	commons.SendResponse(writer, http.StatusOK, json)
}

func GetByID(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	defer db.Close()

	db.Find(&persona, id)

	if persona.ID > 0 {
		json, _ := json.Marshal(persona)
		commons.SendResponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}

func Save(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&persona)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&persona).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(persona)

	commons.SendResponse(writer, http.StatusCreated, json)
}

func Delete(writer http.ResponseWriter, request *http.Request) {

	persona := models.Persona{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&persona, id)

	if persona.ID > 0 {
		db.Delete(persona)
		commons.SendResponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Limita el tamaño del archivo que se puede subir a 10 MB
	r.ParseMultipartForm(10 << 20)

	// Obtén el archivo desde el formulario
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error al obtener el archivo:", err)
		http.Error(w, "No se pudo obtener el archivo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Crea un archivo en el servidor para guardar el archivo subido
	dst, err := os.Create("./uploads/" + handler.Filename)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		http.Error(w, "No se pudo crear el archivo en el servidor", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copia el contenido del archivo subido al archivo en el servidor
	if _, err := io.Copy(dst, file); err != nil {
		fmt.Println("Error al copiar el archivo:", err)
		http.Error(w, "No se pudo copiar el archivo", http.StatusInternalServerError)
		return
	}

	// Responde al cliente que el archivo se ha subido exitosamente
	fmt.Fprintf(w, "Archivo subido exitosamente: %s\n", handler.Filename)
}
