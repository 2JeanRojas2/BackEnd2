package store

import "parcialfinal/internal/domain"

type StoreInterface interface {
	GetOne(id int) (*domain.Dentista, error)
	Create(dentista domain.Dentista) (*domain.Dentista, error)
	Exist(codeValue string) bool
	UpdateOne(dentista domain.Dentista) error
	DeleteOne(id int) error
}
