package dentista

import ("FinalBack/internal/domain"
		"fmt")

type DentistService struct {
	repo DentistRepository
}

// NewDentistService crea un nuevo servicio de dentistas
func NewDentistService(repo DentistRepository) *DentistService {
	return &DentistService{repo: repo}
}

// CreateDentist crea un nuevo dentista
func (s *DentistService) CreateDentist(dentist *domain.Dentista) error {
    err := s.repo.Create(dentist)
    if err != nil {
        return fmt.Errorf("could not create dentist: %v", err)
    }
    return nil
}

// GetDentistByID busca un dentista por ID
func (s *DentistService) GetDentistaById(id int) (*domain.Dentista, error) {
	dentista, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return dentista, nil
}

func (s *DentistService) GetDentistaByMatricula(matricula string) (*domain.Dentista, error) {
	dentista, err := s.repo.GetByMatricula(matricula)
	if err != nil {
		return nil, err
	}
	return dentista, nil
}

// UpdateDentist actualiza un dentista existente
func (s *DentistService) UpdateDentist(dentist *domain.Dentista) error {
    // Check if the dentist exists
    _, err := s.repo.GetById(dentist.Id)
    if err != nil {
        return fmt.Errorf("could not update dentist: %v", err)
    }

    // Update the dentist
    err = s.repo.Update(dentist)
    if err != nil {
        return fmt.Errorf("could not update dentist: %v", err)
    }

    return nil
}

// UpdateDentistField actualiza un campo específico de un dentista existente
func (s *DentistService) UpdateDentistField(id int, field string, value interface{}) error {
	err := s.repo.UpdateDentistField(id, field, value)
	if err != nil {
		return err
	}
	return nil
}

// DeleteDentist elimina un dentista existente
func (s *DentistService) DeleteDentist(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
