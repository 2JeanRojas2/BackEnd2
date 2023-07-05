package turno

import "FinalBack/internal/domain"

type TurnoRepository interface {
	Create(turno *domain.TurnoPost) error
	GetById(id int) (*domain.Turno, error)
	Update(turno *domain.TurnoPost) error
	UpdateField(id int, field string, value interface{}) error
	Delete(id int) error
	CreateByDniAndMatricula(turno *domain.Turno) (*domain.Turno, error)
	GetTurnoByDniPaciente(dni string) ([]*domain.Turno, error)
}