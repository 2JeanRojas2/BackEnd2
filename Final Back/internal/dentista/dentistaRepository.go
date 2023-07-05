package dentista

import "FinalBack/internal/domain"

// DentistRepository define la interfaz para las operaciones con dentistas
type DentistRepository interface {
	Create(dentist *domain.Dentista) error
	GetById(id int) (*domain.Dentista, error)
	GetByMatricula(matricula string) (*domain.Dentista, error)
	Update(dentist *domain.Dentista) error
	UpdateDentistField(dentistID int, field string, value interface{}) error
	Delete(id int) error
    Exist(codeValue string) bool
}
