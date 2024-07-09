package models

import "time"

// Interview struct represents the Interviews table
type Interview struct {
	ID              int       `json:"id"`
	CandidatoID     int       `json:"candidato_id"`
	PuestoID        int       `json:"puesto_id"`
	FechaHora       time.Time `json:"fecha_hora"`
	EntrevistadorID int       `json:"entrevistador_id"`
	Notas           string    `json:"notas"`
	Resultado       string    `json:"resultado"`
}
