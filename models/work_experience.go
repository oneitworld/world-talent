package models

import "time"

// WorkExperience struct represents the WorkExperience table
type WorkExperience struct {
	ID          int       `json:"id"`
	CandidatoID int       `json:"candidato_id"`
	Empresa     string    `json:"empresa"`
	Cargo       string    `json:"cargo"`
	FechaInicio time.Time `json:"fecha_inicio"`
	FechaFin    time.Time `json:"fecha_fin"`
	Descripcion string    `json:"descripcion"`
}
