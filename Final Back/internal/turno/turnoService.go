package turno

import (
	"FinalBack/pkg/web"
	//"FinalBack/internal/dentista"
	"FinalBack/internal/domain"
)

type TurnoService struct {
	turnoRepo TurnoRepository
}

func NewTurnoService(turnoRepo TurnoRepository) *TurnoService {
	return &TurnoService{turnoRepo: turnoRepo}
}

func (s *TurnoService) Create(turno *domain.TurnoPost) error {
	err := s.turnoRepo.Create(turno)
	if err != nil {
		return web.NewBadRequestApiError(err.Error())
	}
	return nil
}

func (s *TurnoService) GetById(id int) (*domain.Turno, error) {
	turno, err := s.turnoRepo.GetById(id)
	if err != nil {
		return nil, web.NewNotFoundApiError(err.Error())
	}
	return turno, nil
}

func (s *TurnoService) Update(turno *domain.TurnoPost) error {
	err := s.turnoRepo.Update(turno)
	if err != nil {
		return web.NewBadRequestApiError(err.Error())
	}
	return nil
}

func (s *TurnoService) UpdateTurnoField(id int, field string, value interface{}) error {
	err := s.turnoRepo.UpdateField(id, field, value)
	if err != nil {
		return web.NewBadRequestApiError(err.Error())
	}
	return nil
}

func (s *TurnoService) Delete(id int) error {
	err := s.turnoRepo.Delete(id)
	if err != nil {
		return web.NewBadRequestApiError(err.Error())
	}
	return nil
}

func (s *TurnoService) CreateTurnoByDniAndMatricula(turno *domain.Turno) (*domain.Turno, error) {
	createdTurno, err := s.turnoRepo.CreateByDniAndMatricula(turno)
	if err != nil {
        return nil, err
    }
	return createdTurno, nil
}

func (s *TurnoService) GetTurnoByDniPaciente(dni string) ([]*domain.Turno, error) {
	turnos, err := s.turnoRepo.GetTurnoByDniPaciente(dni)
	if err != nil {
		return nil, web.NewNotFoundApiError(err.Error())
	}
	return turnos, nil
}