package domain

type Paciente struct {
	Id          int     `json:"id"`
	Nombre      string 	`json:"nombre" binding:"required"`
	Apellido    string 	`json:"apellido" binding:"required"`
	Domicilio 	string 	`json:"domicilio"`
	Dni  		string 	`json:"dni" binding:"required"`
	FechaDeAlta string 	`json:"fechaDeAlta" binding:"required"`
}
