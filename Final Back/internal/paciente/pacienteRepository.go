package paciente

import "FinalBack/internal/domain"

// PacienteRepository define la interfaz para las operaciones con pacientes
type PacienteRepository interface {
	Create(paciente *domain.Paciente) error
	GetById(id int) (*domain.Paciente, error)
	GetPacienteByDNI(dni string) (*domain.Paciente, error)
	Update(paciente *domain.Paciente) error
	UpdatePacienteField(pacienteID int, field string, value interface{}) error
	Delete(id int) error
    Exist(codeValue string) bool
}