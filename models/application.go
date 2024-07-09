package models

import "time"

// Application struct represents the Applications table
type Application struct {
	ID              int       `json:"id"`
	CandidatoID     int       `json:"candidato_id"`
	PuestoID        int       `json:"puesto_id"`
	FechaSolicitud  time.Time `json:"fecha_solicitud"`
	EstadoSolicitud string    `json:"estado_solicitud"`
}
