package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/oneitworld-demo-crud-api-go/commons"
	"github.com/oneitworld-demo-crud-api-go/models"
	"gopkg.in/gomail.v2"
)

const API_NAME = "PERSONA"

type ServiceStatus struct {
	DBStatus string
}

type Email struct {
	EmailTo string
	Subject string
	Content string
}

type APIResponseBody struct {
	Header models.APIResponseHeader
	Data   models.Persona
}

type APIResponseBodyPersonas struct {
	Header models.APIResponseHeader
	Data   []models.Persona
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

// Busca todas la personas en la Base de Datos.
func GetAll(writer http.ResponseWriter, request *http.Request) {
	personas := []models.Persona{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&personas)

	body, err := io.ReadAll(request.Body)
	if err != nil {
		return
	}

	// Convertir el cuerpo a cadena
	bodyString := string(body)

	// Imprimir el cuerpo de la solicitud
	fmt.Println("Cuerpo de la solicitud: ", bodyString)

	// Formatear la fecha y hora con zona horaria
	formattedDateTimeWithZone := time.Now().Format("2006-01-02 15:04:05 MST")

	apiResponseHeader := models.APIResponseHeader{
		Success:   true,
		Message:   "Lista de Personas",
		Datetime:  formattedDateTimeWithZone,
		Channel:   "MOBILE",
		IPAddress: commons.GetIP(request),
	}

	apiResponseBody := APIResponseBodyPersonas{
		Header: apiResponseHeader,
		Data:   personas,
	}

	apiResponseJSON, _ := json.Marshal(apiResponseBody)

	httpStatus := http.StatusOK

	commons.SendResponse(writer, httpStatus, apiResponseJSON)

	// Escribimos la Auditoria en BD
	commons.WriteAudit(request, apiResponseJSON, API_NAME, true, httpStatus, "MOBILE")

}

func GetByID(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	defer db.Close()

	db.Find(&persona, id)

	if persona.ID == 0 {
		commons.SendError(writer, http.StatusNotFound)
		return
	}

	// Formatear la fecha y hora con zona horaria
	formattedDateTimeWithZone := time.Now().Format("2006-01-02 15:04:05 MST")

	apiResponseHeader := models.APIResponseHeader{
		Success:   true,
		Message:   "Persona Encontrada. ID:" + id,
		Datetime:  formattedDateTimeWithZone,
		Channel:   "MOBILE",
		IPAddress: commons.GetIP(request),
	}

	apiResponseBody := APIResponseBody{
		Header: apiResponseHeader,
		Data:   persona,
	}

	apiResponseJSON, _ := json.Marshal(apiResponseBody)

	commons.SendResponse(writer, http.StatusOK, apiResponseJSON)

	// Escribimos la Auditoria en BD
	commons.WriteAudit(request, apiResponseJSON, API_NAME, true, 200, "MOBILE")

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

	go enviarCorreo(persona.Email, "Confirmacion de Registro Persona", "Hola, "+persona.Nombre+"! \nBienvenido a nuestra plataforma de registro de personas! \nONE IT WORLD Team")

	// Formatear la fecha y hora con zona horaria
	formattedDateTimeWithZone := time.Now().Format("2006-01-02 15:04:05 MST")

	apiResponseHeader := models.APIResponseHeader{
		Success:   true,
		Message:   "Persona Creada Exitosamente. ID:" + string(persona.ID),
		Datetime:  formattedDateTimeWithZone,
		Channel:   "MOBILE",
		IPAddress: commons.GetIP(request),
	}

	apiResponseBody := APIResponseBody{
		Header: apiResponseHeader,
		Data:   persona,
	}

	apiResponseJSON, _ := json.Marshal(apiResponseBody)

	commons.SendResponse(writer, http.StatusCreated, apiResponseJSON)

	// Enviar correo de confirmacion

	// Escribimos la Auditoria en BD
	commons.WriteAudit(request, apiResponseJSON, API_NAME, true, http.StatusCreated, "MOBILE")
}

func Delete(writer http.ResponseWriter, request *http.Request) {

	persona := models.Persona{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&persona, id)

	if persona.ID == 0 {
		commons.SendError(writer, http.StatusNotFound)
		return
	}

	// Eliminamos el registro
	db.Delete(persona)

	// Formatear la fecha y hora con zona horaria
	formattedDateTimeWithZone := time.Now().Format("2006-01-02 15:04:05 MST")

	apiResponseHeader := models.APIResponseHeader{
		Success:   true,
		Message:   "Persona [" + id + "] Eliminada Exitosamente",
		Datetime:  formattedDateTimeWithZone,
		Channel:   "MOBILE",
		IPAddress: commons.GetIP(request),
	}

	apiResponseBody := APIResponseBody{
		Header: apiResponseHeader,
		Data:   persona,
	}

	apiResponseJSON, _ := json.Marshal(apiResponseBody)

	commons.SendResponse(writer, http.StatusOK, apiResponseJSON)

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

func SendEmail(writer http.ResponseWriter, request *http.Request) {

	/*
		EMAIL_PWD := os.Getenv("EMAIL_PWD")
		if EMAIL_PWD == "" {
			panic("EMAIL_PWD no está definida")
		}
	*/

	var email Email

	error := json.NewDecoder(request.Body).Decode(&email)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}
	enviarCorreo(email.EmailTo, email.Subject, email.Content)
	// Configuración del mensaje
	/*
		m := gomail.NewMessage()
		m.SetHeader("From", "mnaranjo@oneitworld.com")
		m.SetHeader("To", email.EmailTo)
		m.SetHeader("Subject", "World-Talent-AI: "+email.Subject)
		m.SetBody("text/plain", email.Content)

		// Configuración del servidor SMTP de Gmail
		d := gomail.NewDialer("smtp.gmail.com", 587, "mnaranjo@oneitworld.com", "rtqieswmfxsodnhv")

		// Enviar el correo
		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}

	*/
	fmt.Println("Correo enviado exitosamente!")

	json, _ := json.Marshal(email)

	commons.SendResponse(writer, http.StatusOK, json)

}

func enviarCorreo(emailTo string, subject string, content string) {

	// Configuración del mensaje
	m := gomail.NewMessage()
	m.SetHeader("From", "mnaranjo@oneitworld.com")
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", "World-Talent-AI: "+subject)
	m.SetBody("text/plain", content)

	// Configuración del servidor SMTP de Gmail
	d := gomail.NewDialer("smtp.gmail.com", 587, "mnaranjo@oneitworld.com", "rtqieswmfxsodnhv")

	// Enviar el correo
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("Correo enviado exitosamente!")

}
