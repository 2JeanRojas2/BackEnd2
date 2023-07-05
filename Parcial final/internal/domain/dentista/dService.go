package dentista

import (
	"fmt"
	"parcialfinal/internal/domain"
	"parcialfinal/pkg/web"
)

type IService interface {
	GetDentistaBy(id int) (*domain.Dentista, error)
	CreateDentista(dentista domain.Dentista) (*domain.Dentista, error)
	UpdateDentista(dentista domain.Dentista) error
	DeleteDentistaBy(id int) error
}

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetDentistaBy(id int) (*domain.Dentista, error) {
	return s.repo.GetOne(id)
}

func (s *Service) CreateDentista(dentista domain.Dentista) (*domain.Dentista, error) {
	if s.repo.Exist(dentista.Matricula) {
		return nil, web.NewBadRequestApiError(fmt.Sprintf("dentista with matricula %s already exists", dentista.Matricula))
	}

	return s.repo.Create(dentista)
}

func (s *Service) UpdateDentista(dentista domain.Dentista) error {
	if !s.repo.Exist(dentista.Matricula) {
		return web.NewNotFoundApiError(fmt.Sprintf("dentista with matricula %s not found", dentista.Matricula))
	}

	return s.repo.UpdateOne(dentista)
}

func (s *Service) DeleteDentistaBy(id int) error {
	if err := s.repo.DeleteOne(id); err != nil {
		return web.NewInternalServerApiError(fmt.Sprintf("failed to delete dentista with id %d", id))
	}
	return nil
}
