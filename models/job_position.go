package models

import "time"

// JobPosition struct represents the JobPositions table
type JobPosition struct {
	ID               int       `json:"id"`
	TituloPuesto     string    `json:"titulo_puesto"`
	Descripcion      string    `json:"descripcion"`
	Requisitos       string    `json:"requisitos"`
	FechaPublicacion time.Time `json:"fecha_publicacion"`
	FechaCierre      time.Time `json:"fecha_cierre"`
	Estado           string    `json:"estado"`
}
