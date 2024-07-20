package commons

import (
	"log"
	"net/http"
	"strings"
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

func GetIP(r *http.Request) string {
	// Intentar obtener la IP desde el encabezado X-Forwarded-For (puede contener múltiples IPs)
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// El encabezado X-Forwarded-For puede contener una lista de IPs separadas por comas
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0]) // Devolver la primera IP
	}

	// Intentar obtener la IP desde el encabezado X-Real-IP
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	// Si no se encuentran los encabezados anteriores, usar RemoteAddr
	ip := r.RemoteAddr
	// RemoteAddr también puede contener el puerto, así que lo removemos
	ip = strings.Split(ip, ":")[0]

	return ip
}
