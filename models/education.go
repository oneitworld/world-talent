package models

import "time"

// Education struct represents the Education table
type Education struct {
	ID          int       `json:"id"`
	CandidatoID int       `json:"candidato_id"`
	Institucion string    `json:"institucion"`
	Titulo      string    `json:"titulo"`
	FechaInicio time.Time `json:"fecha_inicio"`
	FechaFin    time.Time `json:"fecha_fin"`
}
