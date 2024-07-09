package models

import "time"

type Candidate struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Apellido        string    `json:"apellido"`
	Email           string    `json:"email"`
	Telefono        string    `json:"telefono"`
	Direccion       string    `json:"direccion"`
	FechaNacimiento time.Time `json:"fecha_nacimiento"`
	Estado          string    `json:"estado"`
}
