package paciente

import (
	"parcialfinal/internal/domain"
	"parcialfinal/pkg/store"
)

type Repository struct {
	Store store.StoreInterface
}

func NewRepository(store store.StoreInterface) *Repository {
	return &Repository{
		Store: store,
	}
}

func (r *Repository) GetPacienteBy(id int) (*domain.Paciente, error) {
	return r.Store.GetOne(id)
}

func (r *Repository) CreatePaciente(paciente domain.Paciente) (*domain.Paciente, error) {
	if r.Store.Exist(paciente.Dni) {
		return nil, domain.ErrDniAlreadyExists
	}
	return r.Store.Create(paciente)
}

func (r *Repository) UpdatePaciente(paciente domain.Paciente) error {
	if r.Store.Exist(paciente.Dni) {
		return domain.ErrDniAlreadyExists
	}
	return r.Store.UpdateOne(paciente)
}

func (r *Repository) DeletePacienteBy(id int) error {
	return r.Store.DeleteOne(id)
}
