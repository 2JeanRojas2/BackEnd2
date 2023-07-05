package paciente

import (
	"FinalBack/internal/domain"
	"fmt"
)

type PacienteService struct {
	repo PacienteRepository
}

// NewPacienteService crea un nuevo servicio de pacientes
func NewPacienteService(repo PacienteRepository) *PacienteService {
	return &PacienteService{repo: repo}
}

// CreatePaciente crea un nuevo paciente
func (s *PacienteService) CreatePaciente(paciente *domain.Paciente) error {
	err := s.repo.Create(paciente)
	if err != nil {
		return fmt.Errorf("could not create patient: %v", err)
	}
	return nil
}

// GetPacienteByID busca un paciente por ID
func (s *PacienteService) GetPacienteById(id int) (*domain.Paciente, error) {
	paciente, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return paciente, nil
}

func (s *PacienteService) GetPacienteByDNI(dni string) (*domain.Paciente, error) {
	paciente, err := s.repo.GetPacienteByDNI(dni)
	if err != nil {
		return nil, err
	}
	return paciente, nil
}

// UpdatePaciente actualiza un paciente existente
func (s *PacienteService) UpdatePaciente(paciente *domain.Paciente) error {
	// Check if the patient exists
	_, err := s.repo.GetById(paciente.Id)
	if err != nil {
		return fmt.Errorf("could not update patient: %v", err)
	}

	// Update the patient
	err = s.repo.Update(paciente)
	if err != nil {
		return fmt.Errorf("could not update patient: %v", err)
	}

	return nil
}

// UpdatePacienteField actualiza un campo específico de un paciente existente
func (s *PacienteService) UpdatePacienteField(id int, field string, value interface{}) error {
	err := s.repo.UpdatePacienteField(id, field, value)
	if err != nil {
		return err
	}
	return nil
}

// DeletePaciente elimina un paciente existente
func (s *PacienteService) DeletePaciente(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}