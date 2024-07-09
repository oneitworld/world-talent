package models

// Interviewer struct represents the Interviewers table
type Interviewer struct {
	ID           int    `json:"id"`
	Nombre       string `json:"nombre"`
	Apellido     string `json:"apellido"`
	Email        string `json:"email"`
	Departamento string `json:"departamento"`
	TituloPuesto string `json:"titulo_puesto"`
}
