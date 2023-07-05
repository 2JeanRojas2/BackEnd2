package domain

type Turno struct {
	Dentista 	Dentista `json:"Dentista" binding:"required"`
	Paciente 	Paciente `json:"Paciente" binding:"required"`
	FechaYHora 	string 	 `json:"FechaYHora" binding:"required"`
	Descripcion string	 `json:"Descripcion" binding:"required"`
}