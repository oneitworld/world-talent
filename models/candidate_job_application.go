package models

import "time"

// CandidateJobApplication struct represents the CandidateJobApplications table
type CandidateJobApplication struct {
	ID              int       `json:"id"`
	CandidatoID     int       `json:"candidato_id"`
	PuestoID        int       `json:"puesto_id"`
	FechaSolicitud  time.Time `json:"fecha_solicitud"`
	EstadoSolicitud string    `json:"estado_solicitud"`
}
