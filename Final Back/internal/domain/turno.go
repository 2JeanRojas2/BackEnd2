package domain

type Turno struct {
	Id          int      `json:"id"`
	Dentista 	Dentista `json:"Dentista" binding:"required"`
	Paciente 	Paciente `json:"Paciente" binding:"required"`
	FechaYHora 	string 	 `json:"FechaYHora" binding:"required"`
	Descripcion string	 `json:"Descripcion" binding:"required"`
}

type TurnoPost struct {
	Id          int      `json:"id"`
	DentistaId 	int `json:"Dentista" binding:"required"`
	PacienteId 	int `json:"Paciente" binding:"required"`
	FechaYHora 	string 	 `json:"FechaYHora" binding:"required"`
	Descripcion string	 `json:"Descripcion" binding:"required"`
}