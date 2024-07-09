package commons

import (
	"log"
	"net/http"
)

func SendResponse(writer http.ResponseWriter, status int, data []byte) {

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(data)

	log.Println(string(data))

}

func SendError(writer http.ResponseWriter, status int) {
	data := []byte(`{}`)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(data)
}
