package dentista

import (
	"fmt"
	"parcialfinal/internal/domain"
	"parcialfinal/pkg/store"
	"parcialfinal/pkg/web"
)

type IRepository interface{
	GetOne(id int) (*domain.Dentista, error)
	Create(dentista domain.Dentista) (*domain.Dentista, error)
	Exist(codeValue string) bool
	UpdateOne(dentista domain.Dentista) error
	DeleteOne(id int) error
}

type Repository struct {
	Storage store.StoreInterface
}

func (r *Repository) GetOne(id int) (*domain.Dentista, error) {
	dentist, err := r.Storage.GetOne(id)
	if err != nil {
		return nil, web.NewNotFoundApiError(fmt.Sprintf("dentista_id %d not found", id))
	}
	return dentist, nil
}

func (r *Repository) Create(dentista domain.Dentista) (*domain.Dentista, error) {
	if r.Storage.Exist(dentista.Matricula) {
		return nil, web.NewBadRequestApiError(fmt.Sprintf("dentista with matricula %s already exists", dentista.Matricula))
	}
	result, err := r.Storage.Create(dentista)
	if err != nil {
		return nil, web.NewInternalServerApiError("failed to create new dentista")
	}
	return result, nil
}

func (r *Repository) UpdateOne(dentista domain.Dentista) error {
	if !r.Storage.Exist(dentista.Matricula) {
		return web.NewNotFoundApiError(fmt.Sprintf("dentista with matricula %s not found", dentista.Matricula))
	}
	err := r.Storage.UpdateOne(dentista)
	if err != nil {
		return web.NewInternalServerApiError(fmt.Sprintf("failed to update dentista with matricula %s", dentista.Matricula))
	}
	return nil
}

func (r *Repository) DeleteOne(id int) error {
	err := r.Storage.DeleteOne(id)
	if err != nil {
		return web.NewInternalServerApiError(fmt.Sprintf("failed to delete dentista with id %d", id))
	}
	return nil
}

