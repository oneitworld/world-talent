package models

import "time"

// Resume struct represents the Resumes table
type Resume struct {
	ID            int       `json:"id"`
	CandidatoID   int       `json:"candidato_id"`
	NombreArchivo string    `json:"nombre_archivo"`
	TipoArchivo   string    `json:"tipo_archivo"`
	URLArchivo    string    `json:"url_archivo"`
	FechaSubida   time.Time `json:"fecha_subida"`
}
