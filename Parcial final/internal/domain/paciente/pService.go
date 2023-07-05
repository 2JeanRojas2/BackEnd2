package paciente

import ("parcialfinal/internal/domain"
		"parcialfinal/pkg/store"
)

type IService interface {
	GetPacienteById(id int) (*domain.Paciente, error)
	CreatePaciente(paciente domain.Paciente) (*domain.Paciente, error)
	UpdatePaciente(paciente domain.Paciente) error
	DeletePacienteById(id int) error
}

type Service struct {
	PacienteStorage Repository
}

func NewPacienteService(pacienteStorage Repository) *Service {
	return &Service{
		PacienteStorage: pacienteStorage,
	}
}

func (s *Service) GetPacienteById(id int) (*domain.Paciente, error) {
	return s.PacienteStorage.GetOne(id)
}

func (s *Service) CreatePaciente(paciente domain.Paciente) (*domain.Paciente, error) {
	return s.PacienteStorage.Create(paciente)
}

func (s *Service) UpdatePaciente(paciente domain.Paciente) error {
	return s.PacienteStorage.UpdateOne(paciente)
}

func (s *Service) DeletePacienteById(id int) error {
	return s.PacienteStorage.DeleteOne(id)
}